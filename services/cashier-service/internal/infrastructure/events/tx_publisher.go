package events

import (
	"context"
	"encoding/json"

	"github.com/studysoros/the-casino-company/services/cashier-service/internal/domain"
	"github.com/studysoros/the-casino-company/shared/contracts"
	"github.com/studysoros/the-casino-company/shared/messaging"
)

type TxEventPublisher struct {
	rabbitmq *messaging.RabbitMQ
}

func NewTxEventPublisher(rabbitmq *messaging.RabbitMQ) *TxEventPublisher {
	return &TxEventPublisher{
		rabbitmq: rabbitmq,
	}
}

func (p *TxEventPublisher) PublishTxDeposited(ctx context.Context, tx *domain.TxModel) error {
	payload := messaging.TxEventData{
		Tx: tx.ToProto(),
	}

	txEventJSON, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	return p.rabbitmq.PublishMessage(ctx, contracts.TxEventDeposited, contracts.AmqpMessage{
		OwnerID: tx.UserId,
		Data:    txEventJSON,
	})
}
