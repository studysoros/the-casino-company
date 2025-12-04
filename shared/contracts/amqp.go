package contracts

type AmqpMessage struct {
	OwnerID string `json:"ownerId"`
	Data    []byte `json:"data"`
}

const (
	// Transaction events
	TxEventDeposited = "tx.event.deposited"
	TxEventWithdrawn = "tx.event.withdrawn"

	// Updating Balance events
	BalanceEventUpdatedSuccess = "balance.event.success"
	BalanceEventUpdatedFailed  = "balance.event.failed"
)
