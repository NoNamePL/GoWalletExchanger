package queryerror

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Func that return json for error on bad request
func WrongQuery(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"Status": "BadRequest",
	})
}
