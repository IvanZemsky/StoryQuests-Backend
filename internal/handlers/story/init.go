package handlers

import (
	"stories-backend/internal/domain/story"

	"github.com/gin-gonic/gin"
)

type StoryHandler struct {
	service domain.StoryService
}

func NewStoryHandler(r *gin.Engine, service domain.StoryService) *StoryHandler {
	handler := StoryHandler{service: service}

	r.GET("/stories", handler.Find)
	r.GET("/stories/:id", handler.FindByID)

	return &handler
}
