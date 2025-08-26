package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stories-backend/internal/domain/auth"
)

func (handler *AuthHandler) GetSession(ctx *gin.Context) {
	token, err := ctx.Cookie(domain.COOKIE_TOKEN)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session, err := handler.service.GetSession(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, session)
}
