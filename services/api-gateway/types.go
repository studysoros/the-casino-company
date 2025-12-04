package main

import (
	pbBalance "github.com/studysoros/the-casino-company/shared/proto/balance"
	pb "github.com/studysoros/the-casino-company/shared/proto/cashier"
)

type depositRequest struct {
	UserId string  `json:"userID"`
	Amount float64 `json:"amount"`
}

type getBalanceRequest struct {
	UserId string `json:"userID"`
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
