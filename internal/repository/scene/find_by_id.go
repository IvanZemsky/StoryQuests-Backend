package repository

import (
	"stories-backend/internal/domain/scene"
	"stories-backend/internal/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (repo *sceneRepository) FindByID(id bson.ObjectID) (domain.Scene, error) {
	ctx, cancel := repository.NewRequestTimeoutContext()
	defer cancel()

	var scene domain.Scene

	err := repo.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&scene)
	if err != nil {
		return domain.Scene{}, err
	}

	return scene, nil
}
