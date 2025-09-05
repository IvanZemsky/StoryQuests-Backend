package handlers

import (
	domain "stories-backend/internal/domain/scene"

	"github.com/gin-gonic/gin"
)

type SceneHandler struct {
	service domain.SceneService
}

func NewSceneHandler(r *gin.Engine, service domain.SceneService) *SceneHandler {
	handler := SceneHandler{service: service}

	r.GET("/stories/:id/scenes", handler.FindByStoryID)
	r.PATCH("/stories/:id/scenes/:sceneId/passes", handler.IncrementPasses)
	r.GET("/stories/:id/results/scenes", handler.FindEndScenesByStoryID)

	return &handler
}
