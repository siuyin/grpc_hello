package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/siuyin/dflt"
	pb "github.com/siuyin/grpc_hello/go/hello"
	"google.golang.org/grpc"
)

//go:generate protoc -I ../../.. --go_out=plugins=grpc:../../hello hello.proto
func main() {
	port := dflt.EnvString("PORT", ":8080")
	fmt.Printf("Greeter server starting on port %s\n", port)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

// server has methods SayHello and Send to implement hello.GreeterServer and hello.Greeter_SayHelloServer.
// This is a small example and I've lumped them together.
type server struct {
	grpc.ServerStream
}

// SayHello implements hello.GreeterServer
func (s *server) SayHello(in *pb.HelloRequest, stream pb.Greeter_SayHelloServer) error {
	log.Printf("Received: %v", in.GetName())
	for i := 0; i < 10; i++ {
		if err := stream.Send(&pb.HelloReply{Message: "Hello " + in.GetName() + ": " + time.Now().Format("15:04:05")}); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

// Send implements hello.Greeter_SayHelloServer and uses grpc.ServerStream to send the message.
func (s *server) Send(m *pb.HelloReply) error {
	return s.ServerStream.SendMsg(m)
}
