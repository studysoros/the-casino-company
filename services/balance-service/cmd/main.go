package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/studysoros/the-casino-company/services/balance-service/internal/infrastructure/events"
	"github.com/studysoros/the-casino-company/services/balance-service/internal/infrastructure/grpc"
	"github.com/studysoros/the-casino-company/services/balance-service/internal/infrastructure/repository"
	"github.com/studysoros/the-casino-company/services/balance-service/internal/service"
	"github.com/studysoros/the-casino-company/shared/env"
	"github.com/studysoros/the-casino-company/shared/messaging"

	grpcserver "google.golang.org/grpc"
)

var GRPCAddr = ":9093"

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	inmemRepo := repository.NewInmemRepository()
	svc := service.NewService(inmemRepo)

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

	grpcServer := grpcserver.NewServer()
	grpc.NewGRPCHandler(grpcServer, svc)

	txConsumer := events.NewTxConsumer(rabbitmq, svc)
	go func() {
		if err := txConsumer.Listen(); err != nil {
			log.Fatalf("tx consumer ailed to listen to the message: %v", err)
		}
	}()

	betSettlementConsumer := events.NewBetSettlementConsumer(rabbitmq, svc, inmemRepo)
	go func() {
		if err := betSettlementConsumer.Listen(); err != nil {
			log.Fatalf("bet settlement consumer failed to listen to the message: %v", err)
		}
	}()

	log.Printf("Starting gRPC server balance service on port %s", lis.Addr().String())

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
