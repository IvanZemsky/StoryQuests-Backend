package service

import (
	domain "stories-backend/internal/domain/scene"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (s *sceneService) FindEndScenesByStoryID(storyID bson.ObjectID) ([]domain.Scene, error) {
	return s.repo.FindEndScenesByStoryID(storyID)
}