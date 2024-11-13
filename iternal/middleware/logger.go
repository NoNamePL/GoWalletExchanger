package logger

import (
	"time"

	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()

		c.Next()

		latency := time.Since(t)

		fmt.Println()
	}
}
