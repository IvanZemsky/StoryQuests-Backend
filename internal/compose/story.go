package compose

import (
	"stories-backend/config"
	domain "stories-backend/internal/domain/story"
	handlers "stories-backend/internal/handlers/story"
	"stories-backend/internal/repository/story"
	"stories-backend/internal/service/story"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type StoryModule struct {
	Handler    *handlers.StoryHandler
	Service    domain.StoryService
	Repository domain.StoryRepository
}

func InitStoryModule(client *mongo.Client, config *config.Config, router *gin.Engine) *StoryModule {
	storyRepository := repository.NewStoryRepository(
		client.Database(config.Database.Name),
		client.Database(config.Database.Name).Collection("stories"),
	)
	storyService := service.NewStoryService(storyRepository)
	storyHandler := handlers.NewStoryHandler(router, storyService)

	return &StoryModule{Handler: storyHandler, Service: storyService, Repository: storyRepository}
}
