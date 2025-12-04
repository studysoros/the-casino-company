package grpc

import (
	"github.com/studysoros/the-casino-company/services/balance-service/internal/domain"
	"github.com/studysoros/the-casino-company/services/balance-service/internal/infrastructure/events"
)

type gRPCHandler struct {
	// TODO: protobuf gRPC

	service   domain.BalanceService
	publisher *events.TxConsumer
}

func NewGRPCHandler(service domain.BalanceService, publisher *events.TxConsumer) *gRPCHandler {
	handler := &gRPCHandler{
		service:   service,
		publisher: publisher,
	}

	// TODO: register handler with grpc server
	return handler
}
