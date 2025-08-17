package domain

import "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	ID    string `json:"id"`
	Login string `json:"login"`
	jwt.RegisteredClaims
}