package postgres

import (
	"database/sql"

	pb "github.com/NoNamePL/GoWalletExchanger/api/gw-wallet-exchanger"
	"github.com/gin-gonic/gin"
)

type HandlerDB struct {
	db     *sql.DB
	client *pb.ExchangeServiceClient
}

func (h *HandlerDB) SetDB(db *sql.DB) {
	h.db = db
}

func (h *HandlerDB) SetClient(client *pb.ExchangeServiceClient) {
	h.client = client
}

// exchange implements storage.DataBase.
func (h *HandlerDB) Exchange(ctx *gin.Context) {
	
}

// getBalance implements storage.DataBase.
func (h *HandlerDB) GetBalance(ctx *gin.Context) {
	
}

// login implements storage.DataBase.
func (h *HandlerDB) Login(ctx *gin.Context) {

}

// rates implements storage.DataBase.
func (h *HandlerDB) Rates(ctx *gin.Context) {

}

// register implements storage.DataBase.
func (h *HandlerDB) Register(ctx *gin.Context) {

}

// sendDeposit implements storage.DataBase.
func (h *HandlerDB) SendDeposit(ctx *gin.Context) {

}

// withdraw implements storage.DataBase.
func (h *HandlerDB) Withdraw(ctx *gin.Context) {

}
