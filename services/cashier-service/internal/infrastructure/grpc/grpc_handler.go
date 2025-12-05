package grpc

import (
	"context"
	"log"

	"github.com/studysoros/the-casino-company/services/cashier-service/internal/domain"
	"github.com/studysoros/the-casino-company/services/cashier-service/internal/infrastructure/events"
	pb "github.com/studysoros/the-casino-company/shared/proto/cashier"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type gRPCHandler struct {
	pb.UnimplementedCashierServiceServer

	service   domain.TxService
	publisher *events.TxEventPublisher
}

func NewGRPCHandler(server *grpc.Server, service domain.TxService, publisher *events.TxEventPublisher) *gRPCHandler {
	handler := &gRPCHandler{
		service:   service,
		publisher: publisher,
	}

	pb.RegisterCashierServiceServer(server, handler)
	return handler
}

func (h *gRPCHandler) Deposit(ctx context.Context, req *pb.DepositRequest) (*pb.DepositResponse, error) {
	userID := req.GetUserID()
	amount := req.GetAmount()

	// TODO: Get real USD/THB FX rate from external API
	fx_rate := 33.0
	amountInUsd := amount / fx_rate

	deposit, err := h.service.Deposit(ctx, userID, amountInUsd)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, "failed to get route: %v", err)
	}

	if err := h.publisher.PublishTxDeposited(ctx, deposit); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to publish the tx deposited event: %v", err)
	}

	return &pb.DepositResponse{
		UserID: deposit.UserId,
		Amount: deposit.Amount,
	}, nil
}

// TODO: Withdraw
