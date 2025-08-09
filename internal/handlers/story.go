package handlers

import (
	"errors"
	"net/http"
	"stories-backend/internal/domain/story"
	"stories-backend/pkg/errors"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	ctx.JSON(http.StatusOK, stories)
}

func (handler *StoryHandler) FindByID(ctx *gin.Context) {
	story, err := handler.service.FindByID(ctx.Param("id"))
	if err != nil {
		if errors.Is(err, customErrors.ErrParsingObjectID) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if errors.Is(err, mongo.ErrNoDocuments) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Story not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}

	ctx.JSON(http.StatusOK, story)
}
