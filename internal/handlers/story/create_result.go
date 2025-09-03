package handlers

import (
	"net/http"
	db "stories-backend/pkg/db/mongo"

	"stories-backend/internal/domain/story"
	handlers "stories-backend/internal/handlers/common"

	"github.com/gin-gonic/gin"
)

func (handler *StoryHandler) CreateResult(ctx *gin.Context) {
	var body createResultBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	userID, err := handlers.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	storyID, err := db.ParseObjectID(body.StoryID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sceneID, err := db.ParseObjectID(body.SceneID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := handler.service.SetResult(domain.SetResultDTO{
		UserID:  userID,
		StoryID: storyID,
		SceneID: sceneID,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

type createResultBody struct {
	StoryID string `json:"storyId"`
	SceneID string `json:"sceneId"`
}
