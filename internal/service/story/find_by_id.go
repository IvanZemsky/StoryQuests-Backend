package service

import (
	domain "stories-backend/internal/domain/story"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (service *storyService) FindByID(id bson.ObjectID) (domain.StoryResponse, error) {
	return service.repo.FindByID(id)
}