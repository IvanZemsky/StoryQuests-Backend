package handlers

import (
	"stories-backend/internal/domain/user"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service domain.UserService
}

func NewUserHandler(r *gin.Engine, service domain.UserService) *UserHandler {
	handler := UserHandler{service: service}

	r.GET("/users/:id", handler.FindByID)

	return &handler
}
