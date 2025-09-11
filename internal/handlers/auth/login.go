package handlers

import (
	"errors"
	"net/http"
	domain "stories-backend/internal/domain/auth"
	commonErrors "stories-backend/pkg/errors"

	"github.com/gin-gonic/gin"
)

func (handler *AuthHandler) Login(ctx *gin.Context) {
	var body domain.LoginDTO
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	token, err := handler.service.Login(body)
	if err != nil {
		if errors.Is(err, commonErrors.ErrNotFound) {
			ctx.JSON(http.StatusBadRequest, commonErrors.ErrLoginUserNotFound)
			return
		}
		if errors.Is(err, commonErrors.ErrMismatchedPassword) {
			ctx.JSON(http.StatusBadRequest, commonErrors.ErrMismatchedPassword)
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	setTokenToCookie(ctx, token)
}
