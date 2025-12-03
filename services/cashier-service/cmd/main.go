package main

import (
	"context"
	"log"
	"net"

	"github.com/studysoros/the-casino-company/services/cashier-service/internal/infrastructure/grpc"
	"github.com/studysoros/the-casino-company/services/cashier-service/internal/infrastructure/repository"
	"github.com/studysoros/the-casino-company/services/cashier-service/internal/service"

	grpcserver "google.golang.org/grpc"
)

var GRPCAddr = ":9092"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	inmemRepo := repository.NewInmemRepository()
	svc := service.NewService(inmemRepo)

	lis, err := net.Listen("tcp", GRPCAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpcserver.NewServer()
	grpc.NewGRPCHandler(grpcServer, svc)

	log.Printf("Starting gRPC cashier service server on port %s", lis.Addr().String())

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Printf("failed to serve: %v", err)
			cancel()
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down the server...")
	grpcServer.GracefulStop()
}
