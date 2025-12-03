package grpc

import (
	"context"
	"log"

	"github.com/studysoros/the-casino-company/services/cashier-service/internal/domain"
	pb "github.com/studysoros/the-casino-company/shared/proto/cashier"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type gRPCHandler struct {
	pb.UnimplementedCashierServiceServer

	service domain.TxService
	// TODO: add event publisher (rabbitmq)
}

func NewGRPCHandler(server *grpc.Server, service domain.TxService) *gRPCHandler {
	handler := &gRPCHandler{
		service: service,
		// TODO: add event publisher (rabbitmq)
	}

	pb.RegisterCashierServiceServer(server, handler)
	return handler
}

func (h *gRPCHandler) Deposit(ctx context.Context, req *pb.DepositRequest) (*pb.DepositResponse, error) {
	userID := req.GetUserID()
	amount := req.GetAmount()

	receipt, err := h.service.Deposit(ctx, userID, amount)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, "failed to get route: %v", err)
	}

	balance := 10.0

	return &pb.DepositResponse{
		UserID:  receipt.UserId,
		Balance: balance + receipt.Amount,
	}, nil
}
