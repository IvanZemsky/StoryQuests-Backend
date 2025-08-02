package service

import "stories-backend/internal/domain"

type storyService struct {
	repo domain.StoryRepository
}

func NewStoryService(repo domain.StoryRepository) domain.StoryService {
	return &storyService{repo: repo}
}

func (service *storyService) Find() ([]domain.Story, error) {
	return service.repo.Find()
}
