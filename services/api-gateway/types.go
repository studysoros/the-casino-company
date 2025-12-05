package main

import (
	pbBalance "github.com/studysoros/the-casino-company/shared/proto/balance"
	pbBetting "github.com/studysoros/the-casino-company/shared/proto/betting"
	pb "github.com/studysoros/the-casino-company/shared/proto/cashier"
)

type depositRequest struct {
	UserId string  `json:"userID"`
	Amount float64 `json:"amount"`
}

type getBalanceRequest struct {
	UserId string `json:"userID"`
}

type betRequest struct {
	UserId  string `json:"userID"`
	BetSide string `json:"betSide"`
}

func (d *depositRequest) toProto() *pb.DepositRequest {
	return &pb.DepositRequest{
		UserID: d.UserId,
		Amount: d.Amount,
	}
}

func (b *getBalanceRequest) toProto() *pbBalance.GetBalanceRequest {
	return &pbBalance.GetBalanceRequest{
		UserID: b.UserId,
	}
}

func (b *betRequest) toProto() *pbBetting.BetRequest {
	return &pbBetting.BetRequest{
		UserID:  b.UserId,
		BetSide: b.BetSide,
	}
}
