package handlers

import (
	"net/http"
	"stories-backend/internal/domain/story"

	"github.com/gin-gonic/gin"
)

type StoryHandler struct {
	service domain.StoryService
}

func NewStoryHandler(r *gin.Engine, service domain.StoryService) {
	handler := StoryHandler{service: service}

	r.GET("/stories", handler.Find)
}

func (handler *StoryHandler) Find(ctx *gin.Context) {
	stories, err := handler.service.Find()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, stories)
}
