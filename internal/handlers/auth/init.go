package handlers

import (
	"log"
	"net/http"
	"stories-backend/internal/domain/auth"
	"time"

	"github.com/gin-gonic/gin"
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

func (handler *AuthHandler) Login(ctx *gin.Context) {
	var body domain.LoginDTO
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	token, error := handler.service.Login(body)
	if error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}
	ctx.SetCookie("token", token, int(time.Hour.Seconds()), "/", "localhost", false, true)
}

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

func (handler *AuthHandler) Logout(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", "localhost", false, true)
}

func (handler *AuthHandler) GetSession(ctx *gin.Context) {
	token, err := ctx.Cookie("token")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(token)

	session, err := handler.service.GetSession(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, session)
}
