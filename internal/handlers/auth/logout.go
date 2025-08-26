package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"stories-backend/internal/domain/auth"
)

func (handler *AuthHandler) Logout(ctx *gin.Context) {
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie(domain.COOKIE_TOKEN, "", -1, "/", "", false, true)
}
