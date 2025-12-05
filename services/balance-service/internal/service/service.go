package service

import (
	"context"
	"fmt"

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

func (s *service) ProcessWinningBet(ctx context.Context, userId string, prize float64) error {
	balance, err := s.repo.GetBalance(ctx, userId)
	if err != nil {
		return fmt.Errorf("failed to get balance: %w", err)
	}
	if balance.Balance < 100 {
		return fmt.Errorf("insufficient balance")
	}
	_, err = s.AddBalance(ctx, userId, prize)
	return err
}

func (s *service) ProcessLosingBet(ctx context.Context, userId string, amount float64) error {
	balance, err := s.repo.GetBalance(ctx, userId)
	if err != nil {
		return fmt.Errorf("failed to get balance: %w", err)
	}
	if balance.Balance < 100 {
		return fmt.Errorf("insufficient balance")
	}
	_, err = s.DeductBalance(ctx, userId, amount)
	return err
}
