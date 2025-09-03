package repository

import (
	"stories-backend/internal/domain/story"
	"stories-backend/internal/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (repo *storyRepository) FindResultByUserIDAndStoryID(userID, storyID bson.ObjectID) (domain.StoryResult, error) {
	ctx, cancel := repository.NewRequestTimeoutContext()
	defer cancel()

	var result domain.StoryResult

	err := repo.resultCollection.FindOne(ctx, bson.M{"userId": userID, "storyId": storyID}).Decode(&result)
	if err != nil {
		return domain.StoryResult{}, err
	}

	return result, nil
}
