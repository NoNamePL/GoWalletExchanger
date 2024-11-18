package postgres

import (
	"database/sql"
	"errors"
	"log/slog"
	"net/http"
	"time"

	pb "github.com/NoNamePL/GoWalletExchanger/api/gw-wallet-exchanger"
	"github.com/NoNamePL/GoWalletExchanger/iternal/config"
	userModel "github.com/NoNamePL/GoWalletExchanger/iternal/model/user"
	utils "github.com/NoNamePL/GoWalletExchanger/pkg/utils/GenerateHashPassword"
	"github.com/NoNamePL/GoWalletExchanger/pkg/utils/queryerror"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type HandlerDB struct {
	db     *sql.DB
	client *pb.ExchangeServiceClient
	logger *slog.Logger
	cfg    *config.Config
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

func (h *HandlerDB) SetConfig(cfg *config.Config) {
	h.cfg = cfg
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
	// storedPassword, exists := user[user.Username]

	stmt, err := h.db.Prepare("SELECT password FROM user WHERE username = ($1)")

	if err != nil {
		queryerror.WrongQuery(ctx)
		return
	}

	var password string

	// Excecute query and write result into variable
	err = stmt.QueryRow(user.Username).Scan(&password)
	if errors.Is(err, sql.ErrNoRows) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Status": "not user in db",
		})
		return
	} else if user.Password != user.Password {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Status": "uncorrect password",
		})
		return
	} else if err != nil {
		queryerror.WrongQuery(ctx)
		return
	}

	// if !exists || storedPassword != user.Password {
	// 	ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	// 	return
	// }

	// create JWT
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &userModel.Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(h.cfg.SecretPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "couldn't create token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})

}

// rates implements storage.DataBase.
func (h *HandlerDB) Rates(ctx *gin.Context) {
	pb.
		pb.GetExchangeRates()
}

// register implements storage.DataBase.
func (h *HandlerDB) Register(ctx *gin.Context) {
	var user userModel.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check user if already exist
	// storedPassword, exists := user[user.Username]

	stmt, err := h.db.Prepare("SELECT password FROM user WHERE username = ($1)")

	if err != nil {
		queryerror.WrongQuery(ctx)
		return
	}

	var password string

	// Excecute query and write result into variable
	err = stmt.QueryRow(user.Username).Scan(&password)
	if !errors.Is(err, sql.ErrNoRows) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Status": "user already exsist",
		})
		return
	}

	// Save user
	stmt, err = h.db.Prepare("INSERT INTO user (username,password) VALUES($1,$2);")

	if err != nil {
		queryerror.WrongQuery(ctx)
		return
	}

	// hash password

	hashPass, err := utils.GenerateHashPassword(password)

	if err != nil {
		h.logger.Error("can't hash password")
		queryerror.WrongQuery(ctx)
		return
	}

	_, err = stmt.Exec(user.Username, hashPass)

	ctx.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// sendDeposit implements storage.DataBase.
func (h *HandlerDB) SendDeposit(ctx *gin.Context) {

}

// withdraw implements storage.DataBase.
func (h *HandlerDB) Withdraw(ctx *gin.Context) {

}
