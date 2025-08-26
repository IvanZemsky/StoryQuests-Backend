package handlers

import (
	"errors"

	"github.com/gin-gonic/gin"
	authDomain "stories-backend/internal/domain/auth"
)

func GetAuthClaims(ctx *gin.Context) (authDomain.JWTClaims, error) {
	claims, exists := ctx.Get(authDomain.CTX_AUTH_CLAIMS)
	if !exists || claims == nil {
		return authDomain.JWTClaims{}, errors.New("no auth claims")
	}

	jwtClaims, ok := claims.(authDomain.JWTClaims)
	if !ok {
		return authDomain.JWTClaims{}, errors.New("invalid auth claims type")
	}

	return jwtClaims, nil
}
