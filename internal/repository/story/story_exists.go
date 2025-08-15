package repository

import (
	"stories-backend/internal/domain/story"
	"stories-backend/internal/repository"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)


func (repo *storyRepository) StoryExists(id bson.ObjectID) (bool, error) {
	ctx, cancel := repository.NewRequestTimeoutContext()
	defer cancel()

	var story domain.Story

	err := repo.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&story)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
