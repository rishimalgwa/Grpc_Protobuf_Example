package main

import (
	"context"
	"fmt"
	heartbeat "hello/heartbeat/heartbeat_pb"
	"io"
	"log"
	"math/rand"

	"google.golang.org/grpc"
)

// helper
func generateBPM() int32 {
	bpm := rand.Intn(100)
	return int32(bpm)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// two way
func NormalAbnormalHeartBeat(c heartbeat.HeartBeatServiceClient) {
	stream, err := c.NormalAbnormalHeartBeat(context.Background())
	handleError(err)
	for t := 0; t < 10; t++ {
		newRequest := &heartbeat.NormalAbnormalHeartBeatRequest{

			Bpm: generateBPM(),
		}
		stream.Send(newRequest)
		fmt.Printf("sent %v\n ", newRequest)
	}
	stream.CloseSend()
	for {
		msg, err := stream.Recv()
		handleError(err)
		if err == io.EOF {
			break
		}
		fmt.Printf("recived %v\n", msg)
	}

}

// one way stream (server)
func UserHeartBeatHistory(c heartbeat.HeartBeatServiceClient) {
	newHistoryRequest := heartbeat.HeartBeatHistoryRequest{
		Username: "Rishi",
	}
	res_stream, err := c.UserHeartBeatHistory(context.Background(), &newHistoryRequest)
	handleError(err)
	for {
		msg, err := res_stream.Recv()
		handleError(err)
		if err == io.EOF {
			break
		}
		fmt.Println(msg)
	}

}

// one way stream (client)
func LiveHeartBeat(c heartbeat.HeartBeatServiceClient) {
	stream, err := c.LiveUserHeartBeat(context.Background())
	handleError(err)
	for t := 0; t < 10; t++ {
		newRequest := &heartbeat.LiveHeartBeatRequest{
			Heartbeat: &heartbeat.HeartBeat{
				Bpm:      generateBPM(),
				Username: "Rishi",
			},
		}
		fmt.Println("Request sent ", newRequest)
		stream.Send(newRequest)
	}
	stream.CloseAndRecv()

}

// unary
func UserHeartBeat(c heartbeat.HeartBeatServiceClient) {
	heartbeatRequest := heartbeat.HeartBeatRequest{
		Heartbeat: &heartbeat.HeartBeat{
			Bpm:      75,
			Username: "Rishi",
		},
	}
	c.UserHeartBeat(context.Background(), &heartbeatRequest)
}
func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	handleError(err)
	defer conn.Close()

	c := heartbeat.NewHeartBeatServiceClient(conn)
	//UserHeartBeat(c)
	//LiveHeartBeat(c)
	UserHeartBeatHistory(c)
}
