package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	domain "stories-backend/internal/domain/auth"
)

func setTokenToCookie(ctx *gin.Context, token string) {
	ctx.SetSameSite(http.SameSiteNoneMode)
	ctx.SetCookie(
		domain.COOKIE_TOKEN,
		token,
		int(time.Hour.Seconds()),
		"/",
		"", 
		true,
		true,
	)
}

func removeTokenFromCookie(ctx *gin.Context) {
	ctx.SetSameSite(http.SameSiteNoneMode)
	ctx.SetCookie(
		domain.COOKIE_TOKEN,
		"",
		-1,
		"/",
		"",
		true,
		true,
	)
}
