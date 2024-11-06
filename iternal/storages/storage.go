package storage

import (
	"context"
)

// Entity представляет собой обобщенный интерфейс для сущностей.
type Entity interface{}

// Database интерфейс для работы с базой данных.
type DataBase interface {
	Connect() error
	Close() error
	Create(ctx context.Context, entity Entity) (int, error)
	Read(ctx context.Context, id int) (Entity, error)
	Update(ctx context.Context, entity Entity) error
	Delete(ctx context.Context, id int) error
	GetCurrencyRequest() CurrencyRequest
	GetExchangeRate() ExchangeRateResponse
}

type CurrencyRequest struct {
	fromCurrency string
	toCurrency string
}

type ExchangeRateResponse struct {
	fromCurrency string
	toCurrency string
	rate float64
}

