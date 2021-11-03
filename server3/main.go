package main

import (
	"log"
	"net"
	"time"

	pb "github.com/leopardx602/grpc/hello"
	"google.golang.org/grpc"
)

type Service struct {
}

func (s *Service) SayHello(in *pb.HelloRequest, stream pb.Hello3Service_SayHelloServer) error {
	// receive
	log.Println(in.GetName())

	// send
	for {
		if err := stream.Send(&pb.HelloResponse{Reply: "Hello, I am server"}); err != nil {
			log.Println("send fail", err)
			return err
		}
		time.Sleep(2 * time.Second)
	}
}

func main() {
	addr := "127.0.0.1:6000"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Server listening on", addr)

	grpcServer := grpc.NewServer()

	pb.RegisterHello3ServiceServer(grpcServer, &Service{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
