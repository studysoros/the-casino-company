package domain

import (
	"context"

	pb "github.com/studysoros/the-casino-company/shared/proto/cashier"
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

func (t *TxModel) ToProto() *pb.Tx {
	return &pb.Tx{
		Id:     t.ID.Hex(),
		UserID: t.UserId,
		Type:   string(t.Type),
		Amount: t.Amount,
	}
}

type TxRepository interface {
	CreateTx(ctx context.Context, tx *TxModel) (*TxModel, error)
}

type TxService interface {
	Deposit(ctx context.Context, userId string, amount float64) (*TxModel, error)
	Withdraw(ctx context.Context, userId string, amount float64) (*TxModel, error)
}
