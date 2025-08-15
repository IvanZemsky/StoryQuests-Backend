package handlers

import (
	"net/http"
	db "stories-backend/pkg/db/mongo"

	"github.com/gin-gonic/gin"
)

func (handler *SceneHandler) FindByStoryID(ctx *gin.Context) {
	id, err := db.ParseObjectID(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	scenes, err := handler.service.FindByStoryID(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, scenes)
}