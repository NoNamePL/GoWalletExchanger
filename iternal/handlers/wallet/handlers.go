package walhandlers

import (
	"database/sql"

	pb "github.com/NoNamePL/GoWalletExchanger/api/gw-wallet-exchanger"
	"github.com/gin-gonic/gin"
)

type HandlerDB struct {
	db     *sql.DB
	client pb.ExchangeServiceClient
}

func RegisterHandlers(router *gin.Engine, db *sql.DB, grpcClient *pb.ExchangeServiceClient) {

	h := HandlerDB{
		db:     db,
		client: *grpcClient,
	}

	router.Group("/api/v1")
	go router.POST("/register", h.register)
	go router.POST("/login", h.login)
	go router.GET("/balance")
	go router.POST("/wallet/deposit")
	go router.POST("/wallet/withdraw")
	go router.GET("/exchange/rates")
	go router.POST("/exchange")
}

func (h *HandlerDB) register(ctx *gin.Context) {

}

func (h *HandlerDB) login(ctx *gin.Context) {

	// TODO: Check to user is on DB
	

}
