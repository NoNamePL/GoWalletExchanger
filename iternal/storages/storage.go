package storage

import (
	"database/sql"
	"log/slog"

	pb "github.com/NoNamePL/GoWalletExchanger/api/gw-wallet-exchanger"
	"github.com/gin-gonic/gin"
)

// Entity представляет собой обобщенный интерфейс для сущностей.
type Entity interface{}

// Database интерфейс для работы с базой данных.
// type DataBase interface {
// 	Connect() error
// 	Close() error
// 	Create(ctx context.Context, entity Entity) (int, error)
// 	Read(ctx context.Context, id int) (Entity, error)
// 	Update(ctx context.Context, entity Entity) error
// 	Delete(ctx context.Context, id int) error
// 	GetCurrencyRequest() CurrencyRequest
// 	GetExchangeRate() ExchangeRateResponse
// }

type DataBase interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	GetBalance(ctx *gin.Context)
	SendDeposit(ctx *gin.Context)
	Withdraw(ctx *gin.Context)
	Rates(ctx *gin.Context)
	Exchange(ctx *gin.Context)
	SetDB(db *sql.DB)
	SetClient(client *pb.ExchangeServiceClient)
	SetLogger(logger *slog.Logger)
}

type CurrencyRequest struct {
	fromCurrency string
	toCurrency   string
}

type ExchangeRateResponse struct {
	fromCurrency string
	toCurrency   string
	rate         float64
}
