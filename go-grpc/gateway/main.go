package main

import (
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	expensepb "github.com/Naveen2070/Go-projects/go-grpc/common/api"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	orderClient := expensepb.NewExpenseServiceClient(conn)

	mux := http.NewServeMux()
	handler := NewHandler(orderClient)
	handler.registerRoutes(mux)

	log.Printf("Gateway server started on %s", ":8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("Failed to start gateway server: ", err)
	}

}
