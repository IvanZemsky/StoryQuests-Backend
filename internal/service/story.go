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
	return service.repo.Find(filters)
}

func (service *storyService) FindByID(id string) (domain.Story, error) {
	objID, err := ParseObjectID(id)
	if err != nil {
        return domain.Story{}, err
    }
	return service.repo.FindByID(objID)
}