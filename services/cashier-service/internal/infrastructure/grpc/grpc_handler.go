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

	deposit, err := h.service.Deposit(ctx, userID, amount)
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

func (h *gRPCHandler) Withdraw(ctx context.Context, req *pb.WithdrawRequest) (*pb.WithdrawResponse, error) {
	userID := req.GetUserID()
	amount := req.GetAmount()

	withdrawal, err := h.service.Withdraw(ctx, userID, amount)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, "failed to get route: %v", err)
	}

	if err := h.publisher.PublishTxWithdrawn(ctx, withdrawal); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to publish the tx withdrawn event: %v", err)
	}

	return &pb.WithdrawResponse{
		UserID: withdrawal.UserId,
		Amount: withdrawal.Amount,
	}, nil
}
