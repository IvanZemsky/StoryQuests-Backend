package service

import (
	"stories-backend/internal/domain/scene"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

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
