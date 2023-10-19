package main

import (
	"context"
	"go_product_grpc/grpc/pb"
	"log"
	"net"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedDataServiceServer
	db *gorm.DB
}

type Product struct {
	gorm.Model
	Id   string
	Name string
}

func (s *Server) SaveProduct(ctx context.Context, product *pb.Product) (*pb.SaveResponse, error) {
	newData := Product{Id: product.Id, Name: product.Name}
	s.db.Create(&newData)
	return &pb.SaveResponse{Success: true}, nil
}

func (s *Server) FindProduct(ctx context.Context, query *pb.QueryRequest) (*pb.QueryResponse, error) {
	var results []*pb.Product
	s.db.Where("id = ?", query.Id).Find(&results)

	protobufResults := make([]*pb.Product, len(results))
	for i, d := range results {
		protobufResults[i] = &pb.Product{Id: d.Id, Name: d.Name}
	}

	return &pb.QueryResponse{Product: protobufResults}, nil
}

func main() {
	port := "50051"
	println("Running rpc server on port: " + port)

	network := "tcp"
	address := "localhost:" + port

	listener, err := net.Listen(network, address)
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("products.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to open the database: %v", err)
	}

	db.AutoMigrate(&Product{})

	println("Success automigrate sqllite database")

	grpcServer := grpc.NewServer()

	pb.RegisterDataServiceServer(grpcServer, &Server{db: db})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed server: %v", err)
	}
}
