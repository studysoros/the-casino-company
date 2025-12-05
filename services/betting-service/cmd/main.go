package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/studysoros/the-casino-company/services/betting-service/internal/infrastructure/events"
	"github.com/studysoros/the-casino-company/services/betting-service/internal/infrastructure/grpc"
	"github.com/studysoros/the-casino-company/services/betting-service/internal/service"
	"github.com/studysoros/the-casino-company/shared/env"
	"github.com/studysoros/the-casino-company/shared/messaging"

	// "github.com/studysoros/the-casino-company/shared/env"
	// "github.com/studysoros/the-casino-company/shared/messaging"

	grpcserver "google.golang.org/grpc"
)

var GRPCAddr = ":9094"

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	svc := service.NewService()

	rabbitmqURI := env.GetString("RABBITMQ_URI", "amqp://guest:guest@rabbitmq:5672/")

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
		<-sigCh
		cancel()
	}()

	lis, err := net.Listen("tcp", GRPCAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	rabbitmq, err := messaging.NewRabbitMQ(rabbitmqURI)
	if err != nil {
		log.Fatal(err)
	}
	defer rabbitmq.Close()

	log.Println("Starting RabbitMQ connection")

	publisher := events.NewBetSettlementPublisher(rabbitmq)

	grpcServer := grpcserver.NewServer()
	grpc.NewGRPCHandler(grpcServer, svc, publisher)

	log.Printf("Starting gRPC server betting service on port %s", lis.Addr().String())

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Printf("failed to serve: %v", err)
			cancel()
		}
	}()

	// wait for the shutdown signal
	<-ctx.Done()
	log.Println("Shutting down the server...")
	grpcServer.GracefulStop()
}
