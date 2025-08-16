package compose

import (
	"stories-backend/config"
	sceneDomain "stories-backend/internal/domain/scene"
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
	SceneRepository sceneDomain.SceneRepository
}

func InitStoryModule(client *mongo.Client, config *config.Config, router *gin.Engine) *StoryModule {
	storyRepository := repository.NewStoryRepository(
		client.Database(config.Database.Name),
		client.Database(config.Database.Name).Collection("stories"),
	)
	storyService := service.NewStoryService(storyRepository, nil)
	storyHandler := handlers.NewStoryHandler(router, storyService)

	return &StoryModule{Handler: storyHandler, Service: storyService, Repository: storyRepository}
}

func (m *StoryModule) SetSceneRepository(sceneRepo sceneDomain.SceneRepository) {
	m.SceneRepository = sceneRepo
	m.Service = service.NewStoryService(m.Repository, sceneRepo)
	m.Handler.UpdateService(m.Service)
}
