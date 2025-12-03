package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TxType string

const (
	TxTypeDeposit  TxType = "deposit"
	TxTypeWithdraw TxType = "withdraw"
)

type TxModel struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	UserId string             `bson:"userID"`
	Type   TxType             `bson:"type"`
	Amount float64            `bson:"amount"`
}

type TxRepository interface {
	CreateTx(ctx context.Context, tx *TxModel) (*TxModel, error)
}

type TxService interface {
	Deposit(ctx context.Context, u string, amount float64) (*TxModel, error)
	Withdraw(ctx context.Context, u string, amount float64) (*TxModel, error)
}
