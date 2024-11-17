package walhandlers

import (
	"database/sql"
	"log/slog"

	pb "github.com/NoNamePL/GoWalletExchanger/api/gw-wallet-exchanger"
	storage "github.com/NoNamePL/GoWalletExchanger/iternal/storages"
	"github.com/NoNamePL/GoWalletExchanger/iternal/storages/postgres"
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(router *gin.Engine, db *sql.DB, grpcClient *pb.ExchangeServiceClient, logger *slog.Logger) {

	var h storage.DataBase

	h = &postgres.HandlerDB{}

	h.SetLogger(logger)
	h.SetDB(db)
	h.SetClient(grpcClient)

	router.Group("/api/v1")
	go router.POST("/register", h.Register)
	go router.POST("/login", h.Login)
	go router.GET("/balance", h.GetBalance)
	go router.POST("/wallet/deposit", h.SendDeposit)
	go router.POST("/wallet/withdraw", h.Withdraw)
	go router.GET("/exchange/rates", h.Rates)
	go router.POST("/exchange", h.Exchange)
}
