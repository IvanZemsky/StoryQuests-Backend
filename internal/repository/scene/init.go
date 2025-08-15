package repository

import (
	domain "stories-backend/internal/domain/scene"
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