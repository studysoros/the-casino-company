package service

import (
	"context"

	"github.com/studysoros/the-casino-company/services/cashier-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type service struct {
	repo     domain.TxRepository
	exchange domain.ExchangeRateProvider
}

func NewService(repo domain.TxRepository, exchange domain.ExchangeRateProvider) *service {
	return &service{
		repo:     repo,
		exchange: exchange,
	}
}

func (s *service) Deposit(ctx context.Context, userId string, amount float64) (*domain.TxModel, error) {
	fx_rate, err := s.exchange.GetUSDRate(ctx)
	if err != nil {
		return nil, err
	}
	amountInUsd := amount / fx_rate

	tx := &domain.TxModel{
		ID:     primitive.NewObjectID(),
		UserId: userId,
		Type:   domain.TxTypeDeposit,
		Amount: amountInUsd,
	}
	return s.repo.CreateTx(ctx, tx)
}

func (s *service) Withdraw(ctx context.Context, userId string, amount float64) (*domain.TxModel, error) {
	// TODO: exchange curreny

	tx := &domain.TxModel{
		ID:     primitive.NewObjectID(),
		UserId: userId,
		Type:   domain.TxTypeWithdraw,
		Amount: amount,
	}
	return s.repo.CreateTx(ctx, tx)
}
