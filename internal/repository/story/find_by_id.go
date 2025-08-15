package repository

import (
	"stories-backend/internal/domain/story"
	"stories-backend/internal/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (repo *storyRepository) FindByID(id bson.ObjectID) (domain.Story, error) {
	ctx, cancel := repository.NewRequestTimeoutContext()
	defer cancel()

	var story domain.Story

	err := repo.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&story)
	if err != nil {
		return domain.Story{}, err
	}

	return story, nil
}
