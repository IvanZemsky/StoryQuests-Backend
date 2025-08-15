package service

import (
	"stories-backend/internal/domain/scene"
	storyDomain "stories-backend/internal/domain/story"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
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

func (service *sceneService) FindByStoryID(id bson.ObjectID) ([]domain.Scene, error) {
	exists, err := service.storyRepo.StoryExists(id)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, mongo.ErrNoDocuments
	}
	return service.repo.FindByStoryID(id)
}
