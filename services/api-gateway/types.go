package main

type depositRequest struct {
	UserId string  `json:"userID"`
	Amount float64 `bson:"amount"`
}
