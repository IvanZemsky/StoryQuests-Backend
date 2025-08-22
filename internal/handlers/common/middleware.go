package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	service "stories-backend/internal/service/auth"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("token")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		claims, err := service.ValidateToken(token)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		ctx.Set("AUTH_CLAIMS", claims)

		ctx.Next()
	}
}
