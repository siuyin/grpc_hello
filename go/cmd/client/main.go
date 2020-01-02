package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/siuyin/dflt"
	"google.golang.org/grpc"

	pb "github.com/siuyin/grpc_hello/go/hello"
)

func main() {
	fmt.Println("hello client in grpc")

	// Connect to server.
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	serverAddr := dflt.EnvString("SERVER_ADDR", "localhost:8080")
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatalf("Could not connect to server at %s: %v",
			serverAddr, err)
	}
	defer conn.Close()

	// Create a client.
	client := pb.NewGreeterClient(conn)
	stream, err := client.SayHello(context.TODO(), &pb.HelloRequest{Name: "gerbau"})
	if err != nil {
		log.Fatalf("could not get stream client: %v", err)
	}

	// Receive client stream.
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("stream recv: %v", err)
		}
		fmt.Printf("received: %v\n", msg.GetMessage())
	}

	fmt.Println("Finishing up.")
}
