package handlers

import (
	"net/http"
	authDomain "stories-backend/internal/domain/auth"
	domain "stories-backend/internal/domain/story"
	db "stories-backend/pkg/db/mongo"

	"github.com/gin-gonic/gin"
)

// @Summary Like or unlike a story
// @Description Toggles like status for a story by authenticated user
// @Tags Story
// @Accept json
// @Produce json
// @Param id path string true "Story ID" format(mongoId)
// @Param request body domain.LikeStoryDTO true "Like status"
// @Success 200 {object} domain.LikeStoryResponse "Like status updated successfully"
// @Failure 400 {object} handlers.BaseErrorResponse "Invalid story ID or request body"
// @Failure 401 {object} handlers.BaseErrorResponse "Unauthorized"
// @Failure 500 {object} handlers.BaseErrorResponse "Internal server error"
// @Security ApiKeyAuth
// @Router /stories/{id}/like [patch]
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
