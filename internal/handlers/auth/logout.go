package handlers

import (
	"github.com/gin-gonic/gin"
)

func (handler *AuthHandler) Logout(ctx *gin.Context) {
	removeTokenFromCookie(ctx)
}
