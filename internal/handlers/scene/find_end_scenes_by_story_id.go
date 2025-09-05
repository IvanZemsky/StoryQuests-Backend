package handlers

import (
	"net/http"
	db "stories-backend/pkg/db/mongo"

	"github.com/gin-gonic/gin"
)

func (handler *SceneHandler) FindEndScenesByStoryID(c *gin.Context) {
	storyID, err := db.ParseObjectID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	scenes, err := handler.service.FindEndScenesByStoryID(storyID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, scenes)

}
