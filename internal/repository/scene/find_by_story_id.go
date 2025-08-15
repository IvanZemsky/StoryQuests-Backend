package repository

import (
	domain "stories-backend/internal/domain/scene"
	"stories-backend/internal/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (repo *sceneRepository) FindByStoryID(storyID bson.ObjectID) ([]domain.Scene, error) {
	ctx, cancel := repository.NewRequestTimeoutContext()
	defer cancel()

	var scenes []domain.Scene

	cursor, err := repo.collection.Find(ctx, bson.M{"storyId": storyID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &scenes); err != nil {
		return nil, err
	}

	return scenes, nil
}
