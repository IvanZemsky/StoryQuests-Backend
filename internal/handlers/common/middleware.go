package handlers

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

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

func CORSMiddleware(origin string) gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{origin},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
