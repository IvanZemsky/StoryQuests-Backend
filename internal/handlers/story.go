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
	r.GET("/stories/:id", handler.FindByID)
}

func (handler *StoryHandler) Find(ctx *gin.Context) {
	stories, err := handler.service.Find(domain.StoryFilters{Search: ctx.Query("search")})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, stories)
}

func (handler *StoryHandler) FindByID(ctx *gin.Context) {
	story, err := handler.service.FindByID(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, story)
}
