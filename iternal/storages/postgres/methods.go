package postgres

import (
	"database/sql"
	"errors"
	"log/slog"
	"net/http"
	"time"

	pb "github.com/NoNamePL/GoWalletExchanger/api/gw-wallet-exchanger"
	userModel "github.com/NoNamePL/GoWalletExchanger/iternal/model/user"
	"github.com/NoNamePL/GoWalletExchanger/pkg/utils/queryerror"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type HandlerDB struct {
	db     *sql.DB
	client *pb.ExchangeServiceClient
	logger *slog.Logger
}

func (h *HandlerDB) SetDB(db *sql.DB) {
	h.db = db
}

func (h *HandlerDB) SetLogger(logger *slog.Logger) {
	h.logger = logger
}

func (h *HandlerDB) SetClient(client *pb.ExchangeServiceClient) {
	h.client = client
}

// exchange implements storage.DataBase.
func (h *HandlerDB) Exchange(ctx *gin.Context) {

}

// getBalance implements storage.DataBase.
func (h *HandlerDB) GetBalance(ctx *gin.Context) {
	// Prepere query for select amount from BD
	stmt, err := h.db.Prepare("SELECT amount FROM wallet WHERE valletId = ($1)")

	if err != nil {
		h.logger.Error("can't prepere db query")
		queryerror.WrongQuery(ctx)
		return
	}

	var resBalance string

	ctx.

	// Excecute query and write result into variable
	err = stmt.QueryRow(id).Scan(&resBalance)
	if errors.Is(err, sql.ErrNoRows) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Status": "not row in db",
		})
		return
	} else if err != nil {
		queryerror.WrongQuery(ctx)
		return
	}

	// return Balance
	ctx.JSON(http.StatusOK, gin.H{
		"Balance": resBalance,
	})
}

// login implements storage.DataBase.
func (h *HandlerDB) Login(ctx *gin.Context) {

	var user userModel.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check user
	storedPassword, exists := user[user.Username]

	if !exists || storedPassword != user.Password {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// create JWT
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &userModel.Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SigningString(jwtKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't create token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})

}

// rates implements storage.DataBase.
func (h *HandlerDB) Rates(ctx *gin.Context) {
	pb.GetExchangeRates{}
}

// register implements storage.DataBase.
func (h *HandlerDB) Register(ctx *gin.Context) {
	var user userModel.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check if user already exists
	if _, exist := user[user.Username]; exist {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	// Save user
	user[user.Username] = user.Password
	ctx.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// sendDeposit implements storage.DataBase.
func (h *HandlerDB) SendDeposit(ctx *gin.Context) {

}

// withdraw implements storage.DataBase.
func (h *HandlerDB) Withdraw(ctx *gin.Context) {

}
