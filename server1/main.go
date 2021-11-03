package main

import (
	"context"
	"log"
	"net"

	pb "github.com/leopardx602/grpc/hello"

	"google.golang.org/grpc"
)

type Service struct {
}

func (s *Service) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v %v", in.GetName(), in.GetAge())
	return &pb.HelloResponse{Reply: "Hello, I am server "}, nil
}

func main() {
	addr := "127.0.0.1:6000"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Server listening on", addr)

	gRPCServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(gRPCServer, &Service{})
	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
