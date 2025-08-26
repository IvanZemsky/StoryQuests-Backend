package handlers

import (
	"errors"
	"net/http"
	domain "stories-backend/internal/domain/auth"
	customErrors "stories-backend/pkg/errors"
	"time"

	"github.com/gin-gonic/gin"
)

func (handler *AuthHandler) Register(ctx *gin.Context) {
	var body domain.RegisterDTO
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
	}
	token, err := handler.service.Register(body)
	if err != nil {
		if errors.Is(err, customErrors.ErrUserAlreadyExists) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie(domain.COOKIE_TOKEN, token, int(time.Hour.Seconds()), "/", "", false, true)
}
