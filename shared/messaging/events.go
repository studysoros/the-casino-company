package messaging

import pb "github.com/studysoros/the-casino-company/shared/proto/cashier"

const (
	TxCreatedEventQueue          = "transaction_created_events"
	BalanceUpdatedResponsesQueue = "balance_updated_responses"
	DeadLetterQueue              = "dead_letter_queue"
)

type TxEventData struct {
	Tx *pb.Tx `json:"tx"`
}
