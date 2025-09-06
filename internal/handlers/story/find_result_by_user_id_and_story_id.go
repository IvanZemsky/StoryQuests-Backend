package handlers

import (
	"net/http"
	db "stories-backend/pkg/db/mongo"

	"github.com/gin-gonic/gin"
)

// @Summary Get user result for a story
// @Description Retrieves a specific user's result for a particular story
// @Tags Story
// @Produce json
// @Param id path string true "Story ID" format(mongoId)
// @Param user_id path string true "User ID" format(mongoId)
// @Success 200 {object} domain.GetStoryResult "User result retrieved successfully"
// @Failure 400 {object} handlers.BaseErrorResponse "Invalid ID format"
// @Failure 404 {object} handlers.BaseErrorResponse "Result not found"
// @Failure 500 {object} handlers.BaseErrorResponse "Internal server error"
// @Router /stories/{id}/results/{user_id} [get]
func (handler *StoryHandler) FindResultByUserIDAndStoryID(ctx *gin.Context) {
	userID, err := db.ParseObjectID(ctx.Param("user_id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	storyID, err := db.ParseObjectID(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stories, err := handler.service.FindResultByUserIDAndStoryID(userID, storyID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, stories)
}
