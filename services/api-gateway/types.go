package main

import (
	pb "github.com/studysoros/the-casino-company/shared/proto/cashier"
)

type depositRequest struct {
	UserId string  `json:"userID"`
	Amount float64 `json:"amount"`
}

func (d *depositRequest) toProto() *pb.DepositRequest {
	return &pb.DepositRequest{
		UserID: d.UserId,
		Amount: d.Amount,
	}
}
