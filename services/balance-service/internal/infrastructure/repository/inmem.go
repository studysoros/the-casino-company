package repository

import (
	"context"

	"github.com/studysoros/the-casino-company/services/balance-service/internal/domain"
)

type inmemRepository struct {
	userBalances map[string]float64
}

func NewInmemRepository() *inmemRepository {
	return &inmemRepository{
		userBalances: make(map[string]float64),
	}
}

func (r *inmemRepository) UpdateBalance(ctx context.Context, update *domain.UpdateModel) (*domain.BalanceModel, error) {
	switch update.Type {
	case "add":
		r.userBalances[update.UserID] += update.Amount
	case "deduct":
		// TODO: check if user has account i.e. balance check if r.userBalances[amount.UserID] even exists
		// TODO: check if user has enough balance
		r.userBalances[update.UserID] -= update.Amount
	}

	return &domain.BalanceModel{
		ID:      update.ID,
		UserID:  update.UserID,
		Balance: r.userBalances[update.UserID],
	}, nil
}
