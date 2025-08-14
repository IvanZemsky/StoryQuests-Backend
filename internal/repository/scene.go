package repository

import (
	domain "stories-backend/internal/domain/scene"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type sceneRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewSceneRepository(db *mongo.Database, collection *mongo.Collection) domain.SceneRepository {
	return &sceneRepository{
		db:         db,
		collection: collection,
	}
}

func (repo *sceneRepository) FindByStoryID(storyID bson.ObjectID) ([]domain.Scene, error) {
	ctx, cancel := NewRequestTimeoutContext()
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
