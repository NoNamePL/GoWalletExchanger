package userModel

import "github.com/dgrijalva/jwt-go"

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
