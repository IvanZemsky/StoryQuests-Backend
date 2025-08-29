package handlers

import (
	"log"
	"net/http"
	sceneDomain "stories-backend/internal/domain/scene"
	domain "stories-backend/internal/domain/story"
	db "stories-backend/pkg/db/mongo"

	"github.com/gin-gonic/gin"
)

func (handler *StoryHandler) Create(ctx *gin.Context) {
	var body createStoryBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	storyDTO, err := createCreateStoryDTO(body.StoryInfo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	story, err := handler.service.Create(storyDTO, body.Scenes)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, story)
}

func createCreateStoryDTO(storyInfoFromBody domain.CreateStoryInfoBody) (domain.CreateStoryDTO, error) {
	authorId, err := db.ParseObjectID(storyInfoFromBody.AuthorID)
	if err != nil {
		log.Println(err)
		return domain.CreateStoryDTO{}, err
	}

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
	StoryInfo domain.CreateStoryInfoBody `json:"storyInfo"`
	Scenes    []sceneDomain.CreateSceneDTO `json:"scenes"`
}
