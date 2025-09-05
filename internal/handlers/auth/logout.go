package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *AuthHandler) Logout(ctx *gin.Context) {
	removeTokenFromCookie(ctx)
	ctx.Status(http.StatusOK)
}
