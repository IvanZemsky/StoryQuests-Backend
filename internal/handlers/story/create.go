package handlers

import (
	"net/http"
	authDomain "stories-backend/internal/domain/auth"
	sceneDomain "stories-backend/internal/domain/scene"
	domain "stories-backend/internal/domain/story"
	db "stories-backend/pkg/db/mongo"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (handler *StoryHandler) Create(ctx *gin.Context) {
	var body createStoryBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	userID, err := getUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	storyDTO, err := createCreateStoryDTO(body.StoryInfo, userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	storyID, err := handler.service.Create(storyDTO, body.Scenes)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusOK, CreateStoryResponse{
		StoryID: storyID.Hex(),
	})
}

func getUserID(ctx *gin.Context) (bson.ObjectID, error) {
	claims, exists := ctx.Get(authDomain.CTX_AUTH_CLAIMS)
	if !exists {
		return bson.ObjectID{}, nil
	}

	stringUserID := claims.(authDomain.JWTClaims).ID

	userID, err := db.ParseObjectID(stringUserID)
	if err != nil {
		return bson.ObjectID{}, err
	}

	return userID, nil
}

func createCreateStoryDTO(
	storyInfoFromBody domain.CreateStoryInfoBody,
	authorId bson.ObjectID,
	) (domain.CreateStoryDTO, error) {
	return domain.CreateStoryDTO{
		Name:        storyInfoFromBody.Name,
		Description: storyInfoFromBody.Description,
		AuthorID:    authorId,
		SceneCount:  storyInfoFromBody.SceneCount,
		Img:         storyInfoFromBody.Img,
		Tags:        storyInfoFromBody.Tags,
	}, nil
}

type createStoryBody struct {
	StoryInfo domain.CreateStoryInfoBody   `json:"storyInfo"`
	Scenes    []sceneDomain.CreateSceneDTO `json:"scenes"`
}

type CreateStoryResponse struct {
	StoryID string `json:"storyId"`
}
