package handlers

import (
	"github.com/gin-gonic/gin"
	"stories-backend/internal/domain/auth"
)

type AuthHandler struct {
	service domain.AuthService
}

func NewAuthHandler(r *gin.Engine, service domain.AuthService) *AuthHandler {
	handler := AuthHandler{service: service}

	r.POST("/auth/login", handler.Login)
	r.POST("/auth/register", handler.Register)
	r.POST("/auth/logout", handler.Logout)
	r.GET("/auth/session", handler.GetSession)

	return &handler
}
