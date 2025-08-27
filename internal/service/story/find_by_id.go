package service

import (
	domain "stories-backend/internal/domain/story"
)

func (service *storyService) FindByID(params domain.FindOneStoryParams) (domain.StoryResponse, error) {
	return service.repo.FindByID(params)
}
