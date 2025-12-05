package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/studysoros/the-casino-company/services/cashier-service/internal/infrastructure/events"
	"github.com/studysoros/the-casino-company/services/cashier-service/internal/infrastructure/exchange"
	"github.com/studysoros/the-casino-company/services/cashier-service/internal/infrastructure/grpc"
	"github.com/studysoros/the-casino-company/services/cashier-service/internal/infrastructure/repository"
	"github.com/studysoros/the-casino-company/services/cashier-service/internal/service"
	"github.com/studysoros/the-casino-company/shared/env"
	"github.com/studysoros/the-casino-company/shared/messaging"

	grpcserver "google.golang.org/grpc"
)

var GRPCAddr = ":9092"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	inmemRepo := repository.NewInmemRepository()
	exchangeRate := exchange.NewFixedExchangeRateProvider(32.0)
	svc := service.NewService(inmemRepo, exchangeRate)

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

	publisher := events.NewTxEventPublisher(rabbitmq)

	grpcServer := grpcserver.NewServer()
	grpc.NewGRPCHandler(grpcServer, svc, publisher)

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
