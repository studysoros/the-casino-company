package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/studysoros/the-casino-company/services/balance-service/internal/infrastructure/events"
	"github.com/studysoros/the-casino-company/services/balance-service/internal/infrastructure/repository"
	"github.com/studysoros/the-casino-company/services/balance-service/internal/service"
	"github.com/studysoros/the-casino-company/shared/env"
	"github.com/studysoros/the-casino-company/shared/messaging"
)

var GRPCAddr = ":9093"

func main() {
	_, cancel := context.WithCancel(context.Background())

	inmemRepo := repository.NewInmemRepository()
	svc := service.NewService(inmemRepo)

	rabbitmqURI := env.GetString("RABBITMQ_URI", "amqp://guest:guest@rabbitmq:5672/")

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
		<-sigCh
		cancel()
	}()

	// TODO: listen for tcp conn

	rabbitmq, err := messaging.NewRabbitMQ(rabbitmqURI)
	if err != nil {
		log.Fatal(err)
	}
	defer rabbitmq.Close()

	log.Println("Starting RabbitMQ connection")

	// TODO: setup gRPC server

	consumer := events.NewTxConsumer(rabbitmq, svc)
	go func() {
		if err := consumer.Listen(); err != nil {
			log.Fatalf("Failed to listen to the message: %v", err)
		}
	}()

	log.Printf("Listening for tx events on port %s", GRPCAddr)

	select {}
}
