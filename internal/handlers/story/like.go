package handlers

import (
	"net/http"
	authDomain "stories-backend/internal/domain/auth"
	domain "stories-backend/internal/domain/story"
	db "stories-backend/pkg/db/mongo"

	"github.com/gin-gonic/gin"
)

func (handler *StoryHandler) LikeStory(ctx *gin.Context) {
	storyID, err := db.ParseObjectID(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claims, exists := ctx.Get(authDomain.CTX_AUTH_CLAIMS)
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	stringUserID := claims.(authDomain.JWTClaims).ID

	userID, err := db.ParseObjectID(stringUserID)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var body struct {
		IsLiked bool `json:"isLiked"`
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
	}

	res, err := handler.service.Like(domain.LikeStoryDTO{
		StoryID: storyID,
		UserID:  userID,
		IsLiked: body.IsLiked,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
