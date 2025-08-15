package service

import (
	"stories-backend/internal/domain/scene"
	storyDomain "stories-backend/internal/domain/story"
)

type sceneService struct {
	repo      domain.SceneRepository
	storyRepo storyDomain.StoryRepository
}

func NewSceneService(
	repo domain.SceneRepository,
	storyRepo storyDomain.StoryRepository,
) domain.SceneService {
	return &sceneService{repo: repo, storyRepo: storyRepo}
}