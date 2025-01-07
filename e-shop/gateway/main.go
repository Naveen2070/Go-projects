package main

import (
	"context"
	"fmt"
	"gateway/handler"
	"log"
	"net/http"
	"time"

	userspb "github.com/Naveen2070/Go-projects/e-shop/common-service/users"
	"github.com/gorilla/mux"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Establish gRPC connection
	conn, err := grpc.NewClient("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Health check before proceeding
	healthClient := userspb.NewHealthServiceClient(conn)
	if err := checkHealth(healthClient); err != nil {
		log.Fatalf("Health check failed: %v", err)
	}

	// Initialize ExpenseService client
	userClient := userspb.NewUserServiceClient(conn)

	// Close the gRPC connection when done
	defer conn.Close()

	// Initialize HTTP server
	mux := mux.NewRouter()

	// Register routes
	handler := handler.NewHandler(userClient)
	handler.RegisterRoutes(mux)

	//helath check
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	log.Printf("Gateway server started on %s", ":8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("Failed to start gateway server: ", err)
	}
}

// checkHealth calls the HealthService Check method
func checkHealth(client userspb.HealthServiceClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.Check(ctx, &userspb.HealthCheckRequest{})
	if err != nil {
		return err
	}

	if resp.Status != "SERVING" {
		return fmt.Errorf("service not serving, status: %s", resp.Status)
	}

	log.Println("Health check successful. Service is SERVING.")
	return nil
}
