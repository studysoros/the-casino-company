package grpc

import (
	"github.com/studysoros/the-casino-company/services/balance-service/internal/domain"
	pb "github.com/studysoros/the-casino-company/shared/proto/balance"

	"google.golang.org/grpc"
)

type gRPCHandler struct {
	pb.UnimplementedBalanceServiceServer

	service domain.BalanceService
}

func NewGRPCHandler(server *grpc.Server, service domain.BalanceService) *gRPCHandler {
	handler := &gRPCHandler{
		service: service,
	}

	pb.RegisterBalanceServiceServer(server, handler)
	return handler
}
