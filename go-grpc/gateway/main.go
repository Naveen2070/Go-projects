package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	expensepb "github.com/Naveen2070/Go-projects/go-grpc/common/api"
	"github.com/Naveen2070/Go-projects/go-grpc/gateway/handler"
)

func main() {
	// Establish gRPC connection
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Health check before proceeding
	healthClient := expensepb.NewHealthServiceClient(conn)
	if err := checkHealth(healthClient); err != nil {
		log.Fatalf("Health check failed: %v", err)
	}

	// Initialize ExpenseService client
	expenseClient := expensepb.NewExpenseServiceClient(conn)

	// Initialize HTTP server
	mux := http.NewServeMux()
	handler := handler.NewHandler(expenseClient)
	handler.RegisterRoutes(mux)

	log.Printf("Gateway server started on %s", ":8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("Failed to start gateway server: ", err)
	}
}

// checkHealth calls the HealthService Check method
func checkHealth(client expensepb.HealthServiceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.Check(ctx, &expensepb.HealthCheckRequest{})
	if err != nil {
		return err
	}

	if resp.Status != "SERVING" {
		return fmt.Errorf("service not serving, status: %s", resp.Status)
	}

	log.Println("Health check successful. Service is SERVING.")
	return nil
}
