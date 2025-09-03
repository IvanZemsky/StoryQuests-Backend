package repository

import (
	"stories-backend/internal/domain/story"
	"stories-backend/internal/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (repo *storyRepository) FindResultsByStoryID(storyID bson.ObjectID) ([]domain.StoryResult, error) {
	ctx, cancel := repository.NewRequestTimeoutContext()
	defer cancel()

	var results []domain.StoryResult

	cursor, err := repo.resultCollection.Find(ctx, bson.M{"storyId": storyID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}
