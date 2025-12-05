package domain

import "context"

type ExchangeRateProvider interface {
	GetUSDRate(ctx context.Context) (float64, error)
}
