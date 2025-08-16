package compose

import (
	domain "stories-backend/internal/domain/scene"
	storyDomain "stories-backend/internal/domain/story"
	handlers "stories-backend/internal/handlers/scene"
	service "stories-backend/internal/service/scene"
)

func InitSceneModule(
	init InitModuleOptions,
	sceneRepository domain.SceneRepository,
	storyRepo storyDomain.StoryRepository,
) {
	sceneService := service.NewSceneService(sceneRepository, storyRepo)
	handlers.NewSceneHandler(init.Router, sceneService)
}
