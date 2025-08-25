package service

import (
	domain "stories-backend/internal/domain/story"
)

func (service *storyService) Find(filters domain.StoryFilters) ([]domain.StoryResponse, int32, error) {
	stories, count, err := service.repo.Find(filters)
	if err != nil {
		return nil, 0, err
	}
	if len(stories) == 0 {
		return []domain.StoryResponse{}, 0, nil
	}
	return stories, count, nil
}
