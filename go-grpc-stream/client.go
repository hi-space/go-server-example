package main

import (
	"context"
	"io"
	"log"

	pb "go-grpc-stream/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	in := &pb.StreamRequest{Id: 1}
	stream, err := client.GetDataStream(context.Background(), in)

	if err != nil {
		log.Fatalf("open stream error %v", err)
	}

	// create a channel
	done := make(chan bool)

	go func() {
		for {
			response, err := stream.Recv()

			if err == io.EOF {
				done <- true // stream is finished
				log.Printf("stream is finished")
				return
			}

			if err != nil {
				log.Fatalf("cannot receive %v", err)
			}

			log.Printf("Received: %s", response.Result)
		}
	}()
	// 위의 go routine이 끝날 때 까지 대기

	<-done //we will wait until all response is received
	log.Printf("finished")
}
