package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/leopardx602/grpc/hello"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":6000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewHello2ServiceClient(conn)
	stream, err := client.SayHello(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// receive
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				log.Println("Done")
				return
			}
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("reply : %v\n", res.GetReply())
		}
	}()

	go func() {
		time.Sleep(8 * time.Second)
		stream.CloseSend()
	}()

	// send
	for {
		if err := stream.Send(&pb.HelloRequest{Name: "Chen"}); err != nil {
			log.Fatal(err)
		}
		time.Sleep(2 * time.Second)
	}

}
