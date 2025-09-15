package handlers

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware(origin string) gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{origin},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
			"Accept",
			"X-Requested-With",
			"Cookie",
		},
		ExposeHeaders: []string{
			"Content-Length",
			"Authorization",
			"X-Total-Count",
			"Set-Cookie",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
