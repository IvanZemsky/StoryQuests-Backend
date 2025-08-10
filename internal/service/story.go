package service

import (
	"stories-backend/internal/domain/story"
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

func (service *storyService) FindByID(id string) (domain.Story, error) {
	objID, err := ParseObjectID(id)
	if err != nil {
		return domain.Story{}, err
	}
	return service.repo.FindByID(objID)
}
