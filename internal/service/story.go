package service

import (
	"stories-backend/internal/domain/story"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type storyService struct {
	repo domain.StoryRepository
}

func NewStoryService(repo domain.StoryRepository) domain.StoryService {
	return &storyService{repo: repo}
}

func (service *storyService) Find(filters domain.StoryFilters) ([]domain.Story, error) {
	stories, err := service.repo.Find(filters)
	if err != nil {
		return nil, err
	}
	if len(stories) == 0 {
		return []domain.Story{}, nil
	}
	return stories, nil
}

func (service *storyService) FindByID(id bson.ObjectID) (domain.Story, error) {
	return service.repo.FindByID(id)
}
