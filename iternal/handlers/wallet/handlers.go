package walhandlers

import (
	"database/sql"

	pb "github.com/NoNamePL/GoWalletExchanger/api/gw-wallet-exchanger"
	storage "github.com/NoNamePL/GoWalletExchanger/iternal/storages"
	"github.com/NoNamePL/GoWalletExchanger/iternal/storages/postgres"
	"github.com/gin-gonic/gin"
)

// type HandlerDB struct {
// 	db     *sql.DB
// 	client pb.ExchangeServiceClient
// }

func RegisterHandlers(router *gin.Engine, db *sql.DB, grpcClient *pb.ExchangeServiceClient) {

	var h storage.DataBase

	h = &postgres.HandlerDB{}
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

// // exchange implements storage.DataBase.
// func (h *HandlerDB) exchange(ctx *gin.Context) {
// 	// register.RegHander(h,ctx)
// }

// // getBalance implements storage.DataBase.
// func (h *HandlerDB) getBalance(ctx *gin.Context) {

// }

// // login implements storage.DataBase.
// func (h *HandlerDB) login(ctx *gin.Context) {

// }

// // rates implements storage.DataBase.
// func (h *HandlerDB) rates(ctx *gin.Context) {

// }

// // register implements storage.DataBase.
// func (h *HandlerDB) register(ctx *gin.Context) {

// }

// // sendDeposit implements storage.DataBase.
// func (h *HandlerDB) sendDeposit(ctx *gin.Context) {

// }

// // withdraw implements storage.DataBase.
// func (h *HandlerDB) withdraw(ctx *gin.Context) {

// }
