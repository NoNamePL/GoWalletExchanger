package logger

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func InitLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
		logger.Debug("Debug message")
		logger.Info("Info message")
		logger.Warn("Warning message")
		logger.Error("Error message")
		return logger
	}
}

func RequestLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		t := time.Now()

		ctx.Next()

		latency := time.Since(t)

		fmt.Println()

	}
}
