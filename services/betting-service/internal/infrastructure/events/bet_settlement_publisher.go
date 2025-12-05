package events

import (
	"context"
	"encoding/json"

	"github.com/studysoros/the-casino-company/services/betting-service/internal/domain"
	"github.com/studysoros/the-casino-company/shared/contracts"
	"github.com/studysoros/the-casino-company/shared/messaging"
)

type BetSettlementPublisher struct {
	rabbitmq *messaging.RabbitMQ
}

func NewBetSettlementPublisher(rabbitmq *messaging.RabbitMQ) *BetSettlementPublisher {
	return &BetSettlementPublisher{
		rabbitmq: rabbitmq,
	}
}

func (p *BetSettlementPublisher) PublishBetSettlementWin(ctx context.Context, b *domain.BetModel) error {
	payload := messaging.BetSettlementEventData{
		BetSettlement: b.ToProto(),
	}

	bsEventJson, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	return p.rabbitmq.PublishMessage(ctx, contracts.BetSettlementWin, contracts.AmqpMessage{
		OwnerID: b.UserId,
		Data:    bsEventJson,
	})
}

func (p *BetSettlementPublisher) PublishBetSettlementLoss(ctx context.Context, b *domain.BetModel) error {
	payload := messaging.BetSettlementEventData{
		BetSettlement: b.ToProto(),
	}

	bsEventJson, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	return p.rabbitmq.PublishMessage(ctx, contracts.BetSettlementLoss, contracts.AmqpMessage{
		OwnerID: b.UserId,
		Data:    bsEventJson,
	})
}
