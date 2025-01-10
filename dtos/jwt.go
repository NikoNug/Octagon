package dtos

import (
	"github.com/golang-jwt/jwt/v4"
)

var JWT_KEY = []byte("octagon")

type JWTClaim struct {
	Username string
	Email    string
	jwt.RegisteredClaims
}
