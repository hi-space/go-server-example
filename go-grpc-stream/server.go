package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	pb "go-grpc-stream/proto"

	"google.golang.org/grpc"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Println(in)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) GetDataStream(in *pb.StreamRequest, srv pb.Greeter_GetDataStreamServer) error {
	log.Printf("response for id : %d", in.Id)

	// use wait group to allow proces to be concurrent
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(count int64) {
			defer wg.Done()

			time.Sleep(time.Duration(count) * time.Second)

			response := pb.StreamResponse{Result: fmt.Sprintf("Request %d for Id: %d", count, in.Id)}

			if err := srv.Send(&response); err != nil {
				log.Printf("send error %v", err)
			}

			log.Printf("Finishing request number: %d", count)
		}(int64(i))
	}

	wg.Wait()

	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
