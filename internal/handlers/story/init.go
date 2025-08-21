package handlers

import (
	"stories-backend/internal/domain/story"
	handlers "stories-backend/internal/handlers/common"

	"github.com/gin-gonic/gin"
)

type StoryHandler struct {
	service domain.StoryService
}

func NewStoryHandler(r *gin.Engine, service domain.StoryService) *StoryHandler {
	handler := StoryHandler{service: service}

	r.GET("/stories", handler.Find)
	r.GET("/stories/:id", handler.FindByID)
	r.PATCH("/stories/:id/like", handlers.AuthMiddleware(), handler.LikeStory)

	return &handler
}
