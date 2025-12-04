package events

import (
	"context"
	"encoding/json"
	"log"

	"github.com/rabbitmq/amqp091-go"
	"github.com/studysoros/the-casino-company/services/balance-service/internal/domain"
	"github.com/studysoros/the-casino-company/shared/contracts"
	"github.com/studysoros/the-casino-company/shared/messaging"
)

type TxConsumer struct {
	rabbitmq *messaging.RabbitMQ
	service  domain.BalanceService
}

func NewTxConsumer(rabbitmq *messaging.RabbitMQ, service domain.BalanceService) *TxConsumer {
	return &TxConsumer{
		rabbitmq: rabbitmq,
		service:  service,
	}
}

func (c *TxConsumer) Listen() error {
	return c.rabbitmq.ConsumeMessages(messaging.TxCreatedEventQueue, func(ctx context.Context, msg amqp091.Delivery) error {
		var txEvent contracts.AmqpMessage
		if err := json.Unmarshal(msg.Body, &txEvent); err != nil {
			log.Printf("Failed to unmarshal message: %v", err)
			return err
		}

		var payload messaging.TxEventData
		if err := json.Unmarshal(txEvent.Data, &payload); err != nil {
			log.Printf("Failed to unmarshal message: %v", err)
			return err
		}

		log.Printf("balance received message: %+v", payload)

		switch msg.RoutingKey {
		case contracts.TxEventDeposited:
			return c.handleDeposit(ctx, payload)
		case contracts.TxEventWithdrawn:
			return c.handleWithdraw(ctx, payload)
		}

		log.Printf("unknown tx event: %+v", payload)

		return nil
	})
}

func (c *TxConsumer) handleDeposit(ctx context.Context, payload messaging.TxEventData) error {
	c.service.AddBalance(ctx, payload.Tx.UserID, payload.Tx.Amount)
	return nil
}

func (c *TxConsumer) handleWithdraw(ctx context.Context, payload messaging.TxEventData) error {
	c.service.DeductBalance(ctx, payload.Tx.UserID, payload.Tx.Amount)
	return nil
}
