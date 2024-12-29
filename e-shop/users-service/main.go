package main

import (
	"log"
	"net"
	"sync"

	userspb "github.com/Naveen2070/Go-projects/e-shop/common-service/users"
	"google.golang.org/grpc"
)

// LoggingListener wraps net.Listener to log connections and disconnections.
type LoggingListener struct {
	net.Listener
	connections sync.Map // To track active connections
}

func (l *LoggingListener) Accept() (net.Conn, error) {
	conn, err := l.Listener.Accept()
	if err != nil {
		return nil, err
	}

	log.Printf("New client connected: %v", conn.RemoteAddr())
	l.connections.Store(conn, struct{}{})

	return &wrappedConn{
		Conn:      conn,
		onCloseFn: func() { l.onConnectionClosed(conn) },
	}, nil
}

func (l *LoggingListener) onConnectionClosed(conn net.Conn) {
	l.connections.Delete(conn)
	log.Printf("Client disconnected: %v", conn.RemoteAddr())
}

type wrappedConn struct {
	net.Conn
	onCloseFn func()
}

func (wc *wrappedConn) Close() error {
	err := wc.Conn.Close()
	wc.onCloseFn()
	return err
}

func main() {
	baseListener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	listener := &LoggingListener{Listener: baseListener}

	grpcServer := grpc.NewServer()
	expenseService := &service.ExpenseServiceServer{}
	healthService := &service.HealthServiceServer{}

	userspb.RegisterExpenseServiceServer(grpcServer, expenseService)
	userspb.RegisterHealthServiceServer(grpcServer, healthService)

	log.Println("gRPC server is running on port :50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
