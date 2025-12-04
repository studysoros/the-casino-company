package service

import (
	"context"

	"github.com/studysoros/the-casino-company/services/balance-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type service struct {
	repo domain.BalanceRepository
}

func NewService(repo domain.BalanceRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) AddBalance(ctx context.Context, userId string, amount float64) (*domain.BalanceModel, error) {
	u := &domain.UpdateModel{
		ID:     primitive.NewObjectID(),
		UserID: userId,
		Type:   domain.UpdateBalanceTypeAdd,
		Amount: amount,
	}
	return s.repo.UpdateBalance(ctx, u)
}

func (s *service) DeductBalance(ctx context.Context, userId string, amount float64) (*domain.BalanceModel, error) {
	u := &domain.UpdateModel{
		ID:     primitive.NewObjectID(),
		UserID: userId,
		Type:   domain.UpdateBalanceTypeDeduct,
		Amount: amount,
	}
	return s.repo.UpdateBalance(ctx, u)
}

func (s *service) GetBalance(ctx context.Context, userId string) (*domain.BalanceModel, error) {
	return s.repo.GetBalance(ctx, userId)
}
