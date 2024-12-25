package main

import (
	"log"
	"net"

	expensepb "github.com/Naveen2070/Go-projects/go-grpc/common/api"
	"github.com/Naveen2070/Go-projects/go-grpc/expense-service/service"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	expenseService := &service.ExpenseServiceServer{}

	expensepb.RegisterExpenseServiceServer(grpcServer, expenseService)

	log.Println("gRPC server is running on port :50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
