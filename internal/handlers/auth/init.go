package handlers

import (
	"stories-backend/internal/domain/auth"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service domain.AuthService
}

func NewAuthHandler(r *gin.Engine, service domain.AuthService) *AuthHandler {
	handler := AuthHandler{service: service}

	r.POST("/auth/login", handler.Login)
	r.POST("/auth/register", handler.Register)
	r.POST("/auth/logout", AuthMiddleware(), handler.Logout)
	r.GET("/auth/session", handler.GetSession)

	return &handler
}
