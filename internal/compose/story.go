package compose

import (
	sceneDomain "stories-backend/internal/domain/scene"
	domain "stories-backend/internal/domain/story"
	handlers "stories-backend/internal/handlers/story"
	"stories-backend/internal/service/story"
)

func InitStoryModule(
	init InitModuleOptions,
	storyRepo domain.StoryRepository,
	sceneRepo sceneDomain.SceneRepository,
) {
	storyService := service.NewStoryService(storyRepo, sceneRepo)
	handlers.NewStoryHandler(init.Router, storyService)
}
