package walhandlers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterHandlers(router *gin.Engine, db *sql.DB) {

	router.Group("/api/v1")
	go router.POST("/register",)
	go router.POST("/login")
	go router.GET("/balance")
	go router.POST("/wallet/deposit")
	go router.POST("/wallet/withdraw")
	go router.GET("/exchange/rates")
	go router.POST("/exchange")
}
