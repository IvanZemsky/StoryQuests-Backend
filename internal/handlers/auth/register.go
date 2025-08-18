package handlers

import (
	"net/http"
	domain "stories-backend/internal/domain/auth"
	"time"

	"github.com/gin-gonic/gin"
)

func (handler *AuthHandler) Register(ctx *gin.Context) {
	var body domain.RegisterDTO
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	token, error := handler.service.Register(body)
	if error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}
	ctx.SetCookie("token", token, int(time.Hour.Seconds()), "/", "localhost", false, true)
}