package grpc

import (
	"context"

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

func (h *gRPCHandler) GetBalance(ctx context.Context, req *pb.GetBalanceRequest) (*pb.GetBalanceResponse, error) {
	userID := req.GetUserID()

	balance, err := h.service.GetBalance(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &pb.GetBalanceResponse{
		Balance: balance.Balance,
	}, nil
}
