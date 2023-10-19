package main

import (
	"context"
	"go_product_grpc/grpc/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedHelloServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {

	return &pb.HelloResponse{Message: "hello world" + in.GetName()}, nil
}

func main() {
	println("Running rpc server")

	network := "tcp"
	address := "localhost:50051"

	listner, err := net.Listen(network, address)

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterHelloServer(grpcServer, &Server{})

	if err := grpcServer.Serve(listner); err != nil {
		log.Fatalf("Failed server: %v", err)
	}

}
