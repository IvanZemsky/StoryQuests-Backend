package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *AuthHandler) Logout(ctx *gin.Context) {
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("token", "", -1, "/", "", false, true)
}