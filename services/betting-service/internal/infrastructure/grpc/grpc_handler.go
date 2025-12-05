package grpc

import (
	"context"

	"github.com/studysoros/the-casino-company/services/betting-service/internal/domain"
	"github.com/studysoros/the-casino-company/services/betting-service/internal/infrastructure/events"
	"github.com/studysoros/the-casino-company/services/betting-service/pkg/types"
	pb "github.com/studysoros/the-casino-company/shared/proto/betting"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type gRPCHandler struct {
	pb.UnimplementedBettingServiceServer

	service   domain.BettingService
	publisher *events.BetSettlementPublisher
}

func NewGRPCHandler(server *grpc.Server, service domain.BettingService, publisher *events.BetSettlementPublisher) *gRPCHandler {
	handler := &gRPCHandler{
		service:   service,
		publisher: publisher,
	}

	pb.RegisterBettingServiceServer(server, handler)
	return handler
}

func (h *gRPCHandler) PlaceBet(ctx context.Context, req *pb.BetRequest) (*pb.BetResponse, error) {
	userID := req.GetUserID()
	betSide := req.GetBetSide()

	bet, err := h.service.PlaceBet(ctx, userID, betSide)
	if err != nil {
		return nil, err
	}

	switch bet.BetResult {
	case types.WinningBet:
		if err := h.publisher.PublishBetSettlementWin(ctx, bet); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to publish the bet settlement win event: %v", err)
		}
	case types.LosingBet:
		if err := h.publisher.PublishBetSettlementLoss(ctx, bet); err != nil {
			return nil, status.Errorf(codes.Internal, "failed to publish the bet settlement loss event: %v", err)
		}
	}

	return &pb.BetResponse{
		BetResult: bet.BetResult,
	}, nil
}
