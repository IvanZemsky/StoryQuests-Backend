package service

import (
	"stories-backend/internal/domain/story"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (s *storyService) FindResultByUserIDAndStoryID(userID, storyID bson.ObjectID) (domain.StoryResult, error) {
	return s.repo.FindResultByUserIDAndStoryID(userID, storyID)
}
