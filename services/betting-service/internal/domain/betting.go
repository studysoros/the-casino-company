package domain

import (
	"context"

	pb "github.com/studysoros/the-casino-company/shared/proto/betting"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BetModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserId    string             `bson:"userID"`
	BetResult string             `bson:"betResult"`
}

func (b *BetModel) ToProto() *pb.BetSettlement {
	return &pb.BetSettlement{
		Id:        b.ID.Hex(),
		UserID:    b.UserId,
		BetResult: b.BetResult,
	}
}

type BettingService interface {
	PlaceBet(ctx context.Context, userId string, betSide string) (*BetModel, error)
}
