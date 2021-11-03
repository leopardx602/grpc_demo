package main

import (
	"context"
	"io"
	"log"

	pb "github.com/leopardx602/grpc/hello"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":6000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewHello3ServiceClient(conn)
	stream, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "Chen"})
	if err != nil {
		log.Fatal(err)
	}

	// receive
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			log.Println("Done")
			return
		}
		if err != nil {
			log.Fatal("receive fail or service down", err)
		}
		log.Printf("reply : %v\n", res.GetReply())
	}
}
