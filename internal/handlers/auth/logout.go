package handlers

import (
	"github.com/gin-gonic/gin"
)

func (handler *AuthHandler) Logout(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", "localhost", false, true)
}