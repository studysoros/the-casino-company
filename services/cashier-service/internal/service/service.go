package service

import (
	"context"

	"github.com/studysoros/the-casino-company/services/cashier-service/internal/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type service struct {
	repo domain.TxRepository
}

func NewService(repo domain.TxRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) Deposit(ctx context.Context, u string, a float64) (*domain.TxModel, error) {
	tx := &domain.TxModel{
		ID:     primitive.NewObjectID(),
		UserId: u,
		Type:   domain.TxTypeDeposit,
		Amount: a,
	}
	return s.repo.CreateTx(ctx, tx)
}

func (s *service) Withdraw(ctx context.Context, u string, a float64) (*domain.TxModel, error) {
	tx := &domain.TxModel{
		ID:     primitive.NewObjectID(),
		UserId: u,
		Type:   domain.TxTypeWithdraw,
		Amount: a,
	}
	return s.repo.CreateTx(ctx, tx)
}
