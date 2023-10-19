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

	client := pb.NewDataServiceClient(con)

	req := &pb.Product{
		Id:   "2",
		Name: "name mock",
	}

	res, err := client.SaveProduct(context.Background(), req)

	if err != nil {
		log.Fatal(err)
	}

	log.Print(res)
}
