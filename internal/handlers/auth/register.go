package handlers

import (
	"errors"
	"net/http"
	domain "stories-backend/internal/domain/auth"
	commonErrors "stories-backend/pkg/errors"

	"github.com/gin-gonic/gin"
)

func (handler *AuthHandler) Register(ctx *gin.Context) {
	var body domain.RegisterDTO
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
	}
	token, err := handler.service.Register(body)
	if err != nil {
		if errors.Is(err, commonErrors.ErrUserAlreadyExists) {
			ctx.JSON(http.StatusBadRequest, commonErrors.ErrUserAlreadyExists)
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	setTokenToCookie(ctx, token)
}
