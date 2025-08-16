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

	return &handler
}

func (h *SceneHandler) UpdateService(newService domain.SceneService) {
	h.service = newService
}
