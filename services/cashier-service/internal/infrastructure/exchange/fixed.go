package exchange

import (
	"context"
)

type FixedExchangeRateProvider struct {
	rate float64
}

func NewFixedExchangeRateProvider(rate float64) *FixedExchangeRateProvider {
	return &FixedExchangeRateProvider{
		rate: rate,
	}
}

func (p *FixedExchangeRateProvider) GetUSDRate(ctx context.Context) (float64, error) {
	return p.rate, nil
}
