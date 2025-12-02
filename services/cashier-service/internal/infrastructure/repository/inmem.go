package repository

import (
	"context"

	"github.com/studysoros/the-casino-company/services/cashier-service/internal/domain"
)

type inmemRepository struct {
	txs map[string]*domain.TxModel
}

func NewInmemRepository() *inmemRepository {
	return &inmemRepository{
		txs: make(map[string]*domain.TxModel),
	}
}

func (r *inmemRepository) CreateTx(ctx context.Context, tx *domain.TxModel) (*domain.TxModel, error) {
	r.txs[tx.ID.Hex()] = tx
	return tx, nil
}
