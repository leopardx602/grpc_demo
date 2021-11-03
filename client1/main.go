package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/leopardx602/grpc/hello"

	"google.golang.org/grpc"
)

func main() {
	addr := "127.0.0.1:6000"
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Can not connect to gRPC server: %v", err)
	}
	defer conn.Close()

	gRPCClient := pb.NewHelloServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := gRPCClient.SayHello(ctx, &pb.HelloRequest{Name: "Chen", Age: 25})
	if err != nil {
		log.Fatalf("Could not get nonce: %v", err)
	}
	fmt.Println("Response:", res.GetReply())
}
