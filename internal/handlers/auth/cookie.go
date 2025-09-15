package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	domain "stories-backend/internal/domain/auth"
)

func setTokenToCookie(ctx *gin.Context, token string) {
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie(domain.COOKIE_TOKEN, token, int(time.Hour.Seconds()), "/", "", false, true)
}

func removeTokenFromCookie(ctx *gin.Context) {
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie(domain.COOKIE_TOKEN, "", -1, "/", "", false, true)
}
