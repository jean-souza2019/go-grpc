package main

import (
	"context"
	"go_product_grpc/protofiles/pb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	con, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewProductServiceClient(con)

	req := &pb.CreateProductRequest{
		Name:        "name product",
		Description: "product teste",
		Price:       75.22,
	}

	res, err := client.CreateProduct(context.Background(), req)

	if err != nil {
		log.Fatal(err)
	}

	log.Print(res)
}
