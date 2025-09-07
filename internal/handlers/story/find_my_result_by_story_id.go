package handlers

import (
	"net/http"
	handlers "stories-backend/internal/handlers/common"
	db "stories-backend/pkg/db/mongo"

	"github.com/gin-gonic/gin"
)

// @Summary Get my result for a story
// @Description Retrieves the authenticated user's result for a specific story
// @Tags stories
// @Produce json
// @Param id path string true "Story ID" format(mongoId)
// @Success 200 {object} domain.GetStoryResult "User result retrieved successfully"
// @Failure 400 {object} handlers.BaseErrorResponse "Invalid story ID format"
// @Failure 401 {object} handlers.BaseErrorResponse "Unauthorized"
// @Failure 404 {object} handlers.BaseErrorResponse "Result not found"
// @Failure 500 {object} handlers.BaseErrorResponse "Internal server error"
// @Security ApiKeyAuth
// @Router /stories/{id}/my-result [get]
func (handler *StoryHandler) FindMyResultByStoryID(ctx *gin.Context) {
	userId, err := handlers.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	storyID, err := db.ParseObjectID(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := handler.service.FindResultByUserIDAndStoryID(userId, storyID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
