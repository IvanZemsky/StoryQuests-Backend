package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"stories-backend/internal/domain/auth"
	"stories-backend/internal/service/auth"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie(domain.COOKIE_TOKEN)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		claims, err := service.ValidateToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		ctx.Set(domain.CTX_AUTH_CLAIMS, claims)

		ctx.Next()
	}
}

func GetSessionMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie(domain.COOKIE_TOKEN)
		if err != nil {
			ctx.Set(domain.CTX_AUTH_CLAIMS, nil)
			ctx.Next()
			return
		}

		claims, err := service.ValidateToken(token)
		if err != nil {
			ctx.Set(domain.CTX_AUTH_CLAIMS, nil)
			ctx.Next()
			return
		}

		ctx.Set(domain.CTX_AUTH_CLAIMS, claims)

		ctx.Next()
	}
}
