package main

import (
	"log"
	"net"
	"sync"

	userspb "github.com/Naveen2070/Go-projects/e-shop/common-service/users"
	"github.com/Naveen2070/Go-projects/e-shop/users-service/service"
	"google.golang.org/grpc"
)

var (
	PORT = ":8081"
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
	baseListener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	listener := &LoggingListener{Listener: baseListener}

	grpcServer := grpc.NewServer()
	userService := &service.UserServiceServer{}
	healthService := &service.HealthServiceServer{}

	userspb.RegisterUserServiceServer(grpcServer, userService)
	userspb.RegisterHealthServiceServer(grpcServer, healthService)

	log.Println("gRPC server is running on port " + PORT + "...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
