package handlers

import (
	"net/http"
	domain "stories-backend/internal/domain/story"
	db "stories-backend/pkg/db/mongo"

	"github.com/gin-gonic/gin"
)

func (handler *StoryHandler) LikeStory(ctx *gin.Context) {
	id, err := db.ParseObjectID(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	var body struct {
		IsLiked bool `json:"isLiked"`
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
	}

	res, err := handler.service.Like(domain.LikeStoryDTO{
		StoryID: id,
		IsLiked: body.IsLiked,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
