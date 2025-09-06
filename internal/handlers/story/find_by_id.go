package handlers

import (
	"errors"
	"log"
	"net/http"
	domain "stories-backend/internal/domain/story"
	handlers "stories-backend/internal/handlers/common"
	db "stories-backend/pkg/db/mongo"
	customErrors "stories-backend/pkg/errors"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// @Summary Get story by ID
// @Description Retrieves a specific story by its unique identifier
// @Tags Story
// @Produce json
// @Param id path string true "Story ID" format(mongoId)
// @Success 200 {object} domain.Story "Story retrieved successfully"
// @Failure 400 {object} handlers.BaseErrorResponse "Invalid ID format or parameters"
// @Failure 404 {object} handlers.BaseErrorResponse "Story not found"
// @Failure 500 {object} handlers.BaseErrorResponse "Internal server error"
// @Security ApiKeyAuth
// @Router /stories/{id} [get]
func (handler *StoryHandler) FindByID(ctx *gin.Context) {
	params := domain.FindOneStoryParams{}

	err := parseParams(ctx, &params)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	story, err := handler.service.FindByID(params)

	if err != nil {
		log.Println(err)
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

func parseParams(ctx *gin.Context, params *domain.FindOneStoryParams) error {
	id, err := db.ParseObjectID(ctx.Param("id"))
	if err != nil {
		return err
	}
	params.ID = id

	authClaims, err := handlers.GetAuthClaims(ctx)
	if err != nil {
		params.Me = bson.NilObjectID
	} else {
		meID, err := db.ParseObjectID(authClaims.ID)
		if err != nil {
			params.Me = bson.NilObjectID
		} else {
			params.Me = meID
		}
	}

	return nil
}
