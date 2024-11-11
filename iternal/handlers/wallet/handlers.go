package walhandlers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	pb "github.com/NoNamePL/GoWalletExchanger/api/gw-wallet-exchanger"
)

type HandlerDB struct {
	db *sql.DB
	client *pb.ExchangeServiceClient
}

func RegisterHandlers(router *gin.Engine, db *sql.DB,grpcClient *pb.ExchangeServiceClient) {

	h := HandlerDB {
		db: db,
		client: grpcClient,
	}

	router.Group("/api/v1")
	go router.POST("/register",)
	go router.POST("/login")
	go router.GET("/balance")
	go router.POST("/wallet/deposit")
	go router.POST("/wallet/withdraw")
	go router.GET("/exchange/rates")
	go router.POST("/exchange")
}

