package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateBalanceType string

const (
	UpdateBalanceTypeAdd    UpdateBalanceType = "add"
	UpdateBalanceTypeDeduct UpdateBalanceType = "deduct"
)

type UpdateModel struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	UserID string             `bson:"userID"`
	Type   UpdateBalanceType  `bson:"type"`
	Amount float64            `bson:"amount"`
}

type BalanceModel struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	UserID  string             `bson:"userID"`
	Balance float64            `bson:"balance"`
}

type BalanceRepository interface {
	UpdateBalance(ctx context.Context, update *UpdateModel) (*BalanceModel, error)
	// TODO: GetBalanceById(ctx context.Context, userId string) (float64, error)
}

type BalanceService interface {
	AddBalance(ctx context.Context, userId string, amount float64) (*BalanceModel, error)
	DeductBalance(ctx context.Context, userId string, amount float64) (*BalanceModel, error)
	// TODO: ShowBalance(ctx context.Context, userId string) (float64, error)
}
