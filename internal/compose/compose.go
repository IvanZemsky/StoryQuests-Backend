package compose

import (
	"stories-backend/config"
	sceneDomain "stories-backend/internal/domain/scene"
	storyDomain "stories-backend/internal/domain/story"
	sceneRepository "stories-backend/internal/repository/scene"
	storyRepository "stories-backend/internal/repository/story"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type repositories struct {
	story storyDomain.StoryRepository
	scene sceneDomain.SceneRepository
}

type InitModuleOptions struct {
	Client *mongo.Client
	Config *config.Config
	Router *gin.Engine
}

func InitModules(init InitModuleOptions) {
	repositories := initRepositories(init.Client, init.Config)

	InitStoryModule(init, repositories.story, repositories.scene)
	InitSceneModule(init, repositories.scene, repositories.story)

}

func initRepositories(
	client *mongo.Client,
	config *config.Config,
) repositories {
	storyRepo := storyRepository.NewStoryRepository(
		client.Database(config.Database.Name),
		client.Database(config.Database.Name).Collection("stories"),
	)

	sceneRepo := sceneRepository.NewSceneRepository(
		client.Database(config.Database.Name),
		client.Database(config.Database.Name).Collection("scenes"),
	)

	return repositories{story: storyRepo, scene: sceneRepo}
}
