package authenticate

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func authenticateMiddleware(ctx *gin.Context) {

	// Retrive the token from cookie
	tokenString, err := ctx.Cookie("token")
	if err != nil {
		ctx.Redirect(http.StatusSeeOther, "/login")
		ctx.Abort()
		return
	}

	// Verify the token
	token, err := verifyToken(tokenString)
	if err != nil {
		ctx.Redirect(http.StatusSeeOther, "/login")
		ctx.Abort()
		return
	}

	// information about the verified token
	fmt.Println(token.Claims)

	// Continue with the next middleware or router handler
	ctx.Next()

}

func verifyToken(tokenString string) (*jwt.Token, error) {
	// Parse the token with the secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	// Check for verification errors
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, err
	}

	return token, nil
}
