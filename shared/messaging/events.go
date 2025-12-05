package messaging

import (
	pbBetting "github.com/studysoros/the-casino-company/shared/proto/betting"
	pb "github.com/studysoros/the-casino-company/shared/proto/cashier"
)

const (
	TxCreatedEventQueue = "transaction_created_events"
	BetSettlementQueue  = "bet_settlement_events"
	DeadLetterQueue     = "dead_letter_queue"
)

type TxEventData struct {
	Tx *pb.Tx `json:"tx"`
}

type BetSettlementEventData struct {
	BetSettlement *pbBetting.BetSettlement `json:"betSettlement"`
}
