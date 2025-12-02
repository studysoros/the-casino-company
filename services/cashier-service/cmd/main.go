package main

import (
	"context"
	"log"
	"time"

	"github.com/studysoros/the-casino-company/services/cashier-service/internal/infrastructure/repository"
	"github.com/studysoros/the-casino-company/services/cashier-service/internal/service"
)

func main() {
	ctx := context.Background()

	inmemRepo := repository.NewInmemRepository()
	svc := service.NewService(inmemRepo)

	t, err := svc.Deposit(ctx, "wangxijao", 100)
	if err != nil {
		log.Println(err)
	}

	log.Println(t)

	for {
		time.Sleep(time.Second)
	}
}
