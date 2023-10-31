package main

import (
	"context"
	example "example/myapp/protos"
	"fmt"
	"log"
	"net"

	grpc "google.golang.org/grpc"
)

type server struct {
	example.UnimplementedExampleServiceServer
}

func (s *server) SayHello(ctx context.Context, req *example.HelloRequest) (*example.HelloResponse, error) {
	return &example.HelloResponse{Greeting: "Hello, " + req.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	example.RegisterExampleServiceServer(grpcServer, &server{})
	fmt.Println("server started")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
