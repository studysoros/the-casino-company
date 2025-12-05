package service

import (
	"context"
	"math/rand"

	"github.com/studysoros/the-casino-company/services/betting-service/internal/domain"
	"github.com/studysoros/the-casino-company/services/betting-service/pkg/types"
)

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) PlaceBet(ctx context.Context, userId string, betSide string) (*domain.BetModel, error) {
	possibleOutcomes := []string{types.SideBlue, types.SideRed}

	randomIndex := rand.Intn(len(possibleOutcomes))
	actualWinningSide := possibleOutcomes[randomIndex]

	var betResult string
	if betSide == actualWinningSide {
		betResult = types.WinningBet
	} else {
		betResult = types.LosingBet
	}

	return &domain.BetModel{
		UserId:    userId,
		BetResult: betResult,
	}, nil
}
