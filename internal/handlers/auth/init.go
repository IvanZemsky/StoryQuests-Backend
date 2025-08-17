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

	r.POST("/login", handler.Login)
	r.POST("/register", handler.Register)
	r.POST("/logout", handler.Logout)
	r.GET("/session", handler.GetSession)

	return &handler
}

func (handler *AuthHandler) Login(ctx *gin.Context) {

}

func (handler *AuthHandler) Register(ctx *gin.Context) {

}

func (handler *AuthHandler) Logout(ctx *gin.Context) {

}

func (handler *AuthHandler) GetSession(ctx *gin.Context){

}
