package service

import (
	sceneDomain "stories-backend/internal/domain/scene"
	"stories-backend/internal/domain/story"
)

type storyService struct {
	repo      domain.StoryRepository
	sceneRepo sceneDomain.SceneRepository
}

func NewStoryService(repo domain.StoryRepository, sceneRepo sceneDomain.SceneRepository) domain.StoryService {
	return &storyService{repo: repo, sceneRepo: sceneRepo}
}
