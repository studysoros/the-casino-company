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

type BetSettlementConsumer struct {
	rabbitmq *messaging.RabbitMQ
	service  domain.BalanceService
}

func NewBetSettlementConsumer(rabbitmq *messaging.RabbitMQ, service domain.BalanceService) *BetSettlementConsumer {
	return &BetSettlementConsumer{
		rabbitmq: rabbitmq,
		service:  service,
	}
}

func (c *BetSettlementConsumer) Listen() error {
	return c.rabbitmq.ConsumeMessages(messaging.BetSettlementQueue, func(ctx context.Context, msg amqp091.Delivery) error {
		var bsEvent contracts.AmqpMessage
		if err := json.Unmarshal(msg.Body, &bsEvent); err != nil {
			log.Printf("Failed to unmarshal message: %v", err)
			return err
		}

		var payload messaging.BetSettlementEventData
		if err := json.Unmarshal(bsEvent.Data, &payload); err != nil {
			log.Printf("Failed to unmarshal message: %v", err)
			return err
		}

		log.Printf("balance received message from BetSettlementQueue: %+v", payload)

		switch msg.RoutingKey {
		case contracts.BetSettlementWin:
			return c.handleWinningBet(ctx, payload)
		case contracts.BetSettlementLoss:
			return c.handleLosingBet(ctx, payload)
		}

		log.Printf("unknown tx event: %+v", payload)

		return nil
	})
}

func (c *BetSettlementConsumer) handleWinningBet(ctx context.Context, payload messaging.BetSettlementEventData) error {
	return c.service.ProcessWinningBet(ctx, payload.BetSettlement.UserID, 90.0)
}

func (c *BetSettlementConsumer) handleLosingBet(ctx context.Context, payload messaging.BetSettlementEventData) error {
	return c.service.ProcessLosingBet(ctx, payload.BetSettlement.UserID, 100.0)
}
