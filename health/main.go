package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc-hello/pkg/health"
	"log"
	"net"
	"os"
	"os/signal"
)

func main() {

	// GRPC Server.
	go func() {
		if err := StartGRPCServer("8080"); err != nil {
			os.Exit(1)
		}
		log.Print("GRPC server stopped")
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	os.Exit(0)

}

// StartGRPCServer sets server's handler.
func StartGRPCServer(port string) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Printf("failed to listen: %v", err)
		return err
	}
	grpcServer := grpc.NewServer()

	hh := health.NewGRPCServer()
	health.RegisterHealthServiceServer(grpcServer, hh)

	log.Printf("GRPC is ready to handle requests %s", lis.Addr().String())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
		return err
	}

	return nil
}
