package main

import (
	"context"
	"fmt"
	heartbeat "hello/heartbeat/heartbeat_pb"
	"io"
	"log"
	"net"
	"os"
	"os/signal"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var collection *mongo.Collection

type server struct {
	heartbeat.UnimplementedHeartBeatServiceServer
}

type heart_item struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Bpm      int32              `bson:"bpm"`
	Username string             `bson:"username"`
}

func pushUserToDb(ctx context.Context, item heart_item) primitive.ObjectID {
	res, err := collection.InsertOne(ctx, item)
	handleError(err)
	return res.InsertedID.(primitive.ObjectID)
}

// two way stream
func (*server) NormalAbnormalHeartBeat(stream heartbeat.HeartBeatService_NormalAbnormalHeartBeatServer) error {
	fmt.Println("NormalAbnormalHeartBeat() called ")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		bpm := req.GetBpm()
		var result string
		if bpm < 60 || bpm > 100 {
			result = fmt.Sprintf("User Heartbeat of %v is Abnormal", bpm)
		} else {
			result = fmt.Sprintf("User Heartbeat of %v is Normal", bpm)

		}
		NAResponse := heartbeat.NormalAbnormalHeartBeatResponse{
			Result: result,
		}
		fmt.Printf("Sending back response%v\n", result)
		stream.Send(&NAResponse)
	}
}

//one way stream (server)
func (*server) UserHeartBeatHistory(req *heartbeat.HeartBeatHistoryRequest, stream heartbeat.HeartBeatService_UserHeartBeatHistoryServer) error {
	fmt.Println("HeartBeatHistory() called")
	username := req.GetUsername()
	filter := bson.M{
		"username": username,
	}
	var result_data []heart_item
	cursor, err := collection.Find(context.TODO(), filter)
	handleError(err)

	cursor.All(context.Background(), &result_data)
	for _, v := range result_data {
		historyResponse := heartbeat.HeartBeatHistoryResponse{
			HeartBeat: &heartbeat.HeartBeat{
				Bpm:      v.Bpm,
				Username: v.Username,
			},
		}
		stream.Send(&historyResponse)
	}
	return nil
}

// one way stream (client)
func (*server) LiveHeartBeat(stream heartbeat.HeartBeatService_LiveUserHeartBeatServer) error {
	result := ""
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&heartbeat.LiveHeartBeatResponse{
				Result: result,
			})
		}
		handleError(err)
		bpm := msg.Heartbeat.GetBpm()
		username := msg.Heartbeat.GetUsername()
		docID := pushUserToDb(context.TODO(), heart_item{
			Bpm:      bpm,
			Username: username,
		})
		result += fmt.Sprintf("User Heartbeat = %v, docid = %v ----", bpm, docID)
	}
}

// Unary
func (*server) UserHeartBeat(ctx context.Context, req *heartbeat.HeartBeatRequest) (*heartbeat.HeartBeatResponse, error) {
	fmt.Println("userheartbeat called!")
	bpm := req.GetHeartbeat().GetBpm()
	username := req.Heartbeat.GetUsername()
	newHeartItem := heart_item{
		Bpm:      bpm,
		Username: username,
	}
	docId := pushUserToDb(ctx, newHeartItem)
	result := fmt.Sprintf("user heartbeat is %v, newly created docid is %v", bpm, docId)
	heartbeatBeatResponse := heartbeat.HeartBeatResponse{
		Result: result,
	}
	return &heartbeatBeatResponse, nil
}
func main() {
	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	heartbeat.RegisterHeartBeatServiceServer(s, &server{})

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	handleError(err)
	go func() {
		if err := s.Serve(lis); err != nil {
			handleError(err)
		}
	}()
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	handleError(err)
	fmt.Println("MongoDB connected")
	err = client.Connect(context.TODO())
	handleError(err)
	collection = client.Database("heartbeat").Collection("heartbeat")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	fmt.Println("closing Mongo Connection")
	if err := client.Disconnect(context.TODO()); err != nil {
		handleError(err)

	}
	s.Stop()
}
