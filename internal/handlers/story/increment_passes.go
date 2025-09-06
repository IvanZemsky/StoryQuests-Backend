package handlers

import (
	"net/http"
	db "stories-backend/pkg/db/mongo"

	"github.com/gin-gonic/gin"
)

// @Summary Increment story passes counter
// @Description Increments the number of passes for a specific story
// @Tags Story
// @Produce json
// @Param id path string true "Story ID" format(mongoId)
// @Success 200 "Passes counter incremented successfully"
// @Failure 400 {object} handlers.BaseErrorResponse "Invalid story ID"
// @Failure 500 {object} handlers.BaseErrorResponse "Internal server error"
// @Router /stories/{id}/passes [patch]
func (handler *StoryHandler) IncrementPasses(ctx *gin.Context) {
	id, err := db.ParseObjectID(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = handler.service.IncrementPasses(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
