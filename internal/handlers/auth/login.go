package handlers

import (
	"net/http"
	"errors"
	domain "stories-backend/internal/domain/auth"
	"stories-backend/pkg/errors"
	"time"

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
		if errors.Is(err, customErrors.ErrLoginUserNotFound) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, customErrors.ErrMismatchedPassword) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// в отдельную функцию
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie(domain.COOKIE_TOKEN, token, int(time.Hour.Seconds()), "/", "", false, true)
}