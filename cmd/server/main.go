package main

import (
	"context"
	"go_product_grpc/protofiles/pb"
	"log"
	"net"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedProductServiceServer
	db *gorm.DB
}

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       float32
}

func (s *Server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	newProduct := Product{Name: req.GetName(), Description: req.GetDescription(), Price: req.GetPrice()}
	s.db.Create(&newProduct)

	return &pb.CreateProductResponse{
		Product: &pb.Product{
			Id:          int32(newProduct.ID),
			Name:        newProduct.Name,
			Description: newProduct.Description,
			Price:       newProduct.Price,
		},
	}, nil
}

func (s *Server) FindProducts(ctx context.Context, req *pb.FindProductsRequest) (*pb.FindProductsResponse, error) {
	var products []Product
	s.db.Find(&products)

	var responseProducts []*pb.Product
	for _, p := range products {
		responseProducts = append(responseProducts, &pb.Product{
			Id:          int32(p.ID),
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
		})
	}

	return &pb.FindProductsResponse{
		Products: responseProducts,
	}, nil
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

	pb.RegisterProductServiceServer(grpcServer, &Server{db: db})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed server: %v", err)
	}
}
