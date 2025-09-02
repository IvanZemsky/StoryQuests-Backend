package repository

import (
	"stories-backend/internal/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (repo *sceneRepository) IncrementPasses(sceneID bson.ObjectID) error {
	ctx, cancel := repository.NewRequestTimeoutContext()
	defer cancel()

	_, err := repo.collection.UpdateOne(
		ctx,
		bson.M{
			"_id": sceneID,
		},
		bson.M{"$inc": bson.M{"passes": 1}})
	if err != nil {
		return err
	}

	return nil
}
