package service

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"stories-backend/internal/domain/scene"
)

func (repo *sceneService) FindInStoryByID(storyID bson.ObjectID, sceneID bson.ObjectID) (domain.Scene, error) {
	return repo.FindInStoryByID(storyID, sceneID)
}
