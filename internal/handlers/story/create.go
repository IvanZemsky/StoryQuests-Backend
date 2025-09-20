package handlers

import (
	"net/http"
	sceneDomain "stories-backend/internal/domain/scene"
	domain "stories-backend/internal/domain/story"
	handlers "stories-backend/internal/handlers/common"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// @Summary Create story
// @Description Create story
// @Tags Story
// @Accept json
// @Produce json
// @Param storyInfo body domain.CreateStoryInfoBody true "Story info"
// @Success 201
// @Failure 400
// @Router /stories [post]
func (handler *StoryHandler) Create(ctx *gin.Context) {
	var body createStoryBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	userID, err := handlers.GetUserID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	storyDTO, err := createCreateStoryDTO(&body.StoryInfo, userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	storyID, err := handler.service.Create(storyDTO, body.Scenes)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	ctx.JSON(http.StatusCreated, CreateStoryResponse{
		StoryID: storyID.Hex(),
	})
}

func createCreateStoryDTO(
	storyInfoFromBody *domain.CreateStoryInfoBody,
	authorId bson.ObjectID,
	) (*domain.CreateStoryDTO, error) {
	return &domain.CreateStoryDTO{
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
