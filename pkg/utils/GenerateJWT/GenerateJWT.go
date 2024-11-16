package generatejwt

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// create a new JWT token with claims
func CreateToken(userName string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userName,
		"iss": "wallet-exchanger",
		"aud":"",
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, err := claims.SignedString(secretKey)
	if err != nil {
		return "",err
	}

	return tokenString, nil
}
