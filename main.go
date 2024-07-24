package main

import (
	"fmt"
	"log"
	"net"

	"github.com/kodinggo/category-service-gp1/internal/delivery/grpcsvc"
	pb "github.com/kodinggo/category-service-gp1/pb/category"
	"google.golang.org/grpc"
)

const grpcPort = 8084

func main() {
	// Run grpc server
	grpcServer := grpc.NewServer()

	categoryService := grpcsvc.NewCategoryService()

	pb.RegisterCategoryServiceServer(grpcServer, categoryService)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Panicf("failed create http listener, error: %v", err)
	}

	log.Printf("grpc server is running with port: :%d", grpcPort)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Panicf("failed run grpc server, error: %v", err)
	}
}
