package main

import (
	"context"
	"go_product_grpc/grpc/pb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	con, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewHelloClient(con)

	req := &pb.HelloRequest{
		Name: "name mock",
	}

	res, err := client.SayHello(context.Background(), req)

	if err != nil {
		log.Fatal(err)
	}

	log.Print(res)
}
