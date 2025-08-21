package service

import (
	domain "stories-backend/internal/domain/story"
)

func (service *storyService) Find(filters domain.StoryFilters) ([]domain.StoryResponse, error) {
	stories, err := service.repo.Find(filters)
	if err != nil {
		return nil, err
	}
	if len(stories) == 0 {
		return []domain.StoryResponse{}, nil
	}
	return stories, nil
}
