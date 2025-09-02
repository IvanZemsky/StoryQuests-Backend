package service

import (
	domain "stories-backend/internal/domain/story"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (s *storyService) IncrementPasses(storyID bson.ObjectID) error {
	_, err := s.repo.FindByID(domain.FindOneStoryParams{ID: storyID, Me: bson.NilObjectID})
	if err != nil {
		return err
	}

	return s.repo.IncrementPasses(storyID)
}
