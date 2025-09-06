package handlers

import (
	"net/http"
	db "stories-backend/pkg/db/mongo"

	"github.com/gin-gonic/gin"
)

// @Summary Get results by story ID
// @Description Retrieves all results for a specific story
// @Tags Story
// @Produce json
// @Param id path string true "Story ID" format(mongoId)
// @Success 200 {array} domain.GetStoryResult "Results retrieved successfully"
// @Failure 400 {object} handlers.BaseErrorResponse "Invalid story ID format"
// @Failure 500 {object} handlers.BaseErrorResponse "Internal server error"
// @Router /stories/{id}/results [get]
func (handler *StoryHandler) FindResultsByStoryID(ctx *gin.Context) {
	storyID, err := db.ParseObjectID(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stories, err := handler.service.FindResultsByStoryID(storyID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, stories)
}
