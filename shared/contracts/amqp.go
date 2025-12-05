package contracts

type AmqpMessage struct {
	OwnerID string `json:"ownerId"`
	Data    []byte `json:"data"`
}

const (
	// Transaction events
	TxEventDeposited = "tx.event.deposited"
	TxEventWithdrawn = "tx.event.withdrawn"

	// Bet events
	BetSettlementWin  = "bet.settlement.win"
	BetSettlementLoss = "bet.settlement.loss"
)
