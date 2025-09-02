package handlers

import (
	"stories-backend/internal/domain/story"
	authHandlers "stories-backend/internal/handlers/auth"

	"github.com/gin-gonic/gin"
)

type StoryHandler struct {
	service domain.StoryService
}

func NewStoryHandler(r *gin.Engine, service domain.StoryService) *StoryHandler {
	handler := StoryHandler{service: service}

	r.GET("/stories", authHandlers.GetSessionMiddleware(), handler.Find)
	r.GET("/stories/:id", authHandlers.GetSessionMiddleware(), handler.FindByID)
	r.PATCH("/stories/:id/like", authHandlers.AuthMiddleware(), handler.LikeStory)
	r.PATCH("/stories/:id/passes", handler.IncrementPasses)
	r.POST("/stories/create", authHandlers.AuthMiddleware(), handler.Create)

	return &handler
}
