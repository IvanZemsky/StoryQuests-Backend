package handlers

import (
	"errors"
	"net/http"
	db "stories-backend/pkg/db/mongo"
	customErrors "stories-backend/pkg/errors"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (handler *StoryHandler) FindByID(ctx *gin.Context) {
	id, err := db.ParseObjectID(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	story, err := handler.service.FindByID(id)

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
		return
	}

	ctx.JSON(http.StatusOK, story)
}