package compose

import (
	"stories-backend/config"
	domain "stories-backend/internal/domain/scene"
	storyDomain "stories-backend/internal/domain/story"
	handlers "stories-backend/internal/handlers/scene"
	"stories-backend/internal/repository/scene"
	"stories-backend/internal/service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SceneModule struct {
	Handler    *handlers.SceneHandler
	Service    domain.SceneService
	Repository domain.SceneRepository
}

func InitSceneModule(
	client *mongo.Client,
	config *config.Config,
	router *gin.Engine,
	storyRepo storyDomain.StoryRepository,
) *SceneModule {
	sceneRepository := repository.NewSceneRepository(
		client.Database(config.Database.Name),
		client.Database(config.Database.Name).Collection("scenes"),
	)
	sceneService := service.NewSceneService(sceneRepository, storyRepo)
	sceneHandler := handlers.NewSceneHandler(router, sceneService)

	return &SceneModule{Handler: sceneHandler, Service: sceneService, Repository: sceneRepository}
}
