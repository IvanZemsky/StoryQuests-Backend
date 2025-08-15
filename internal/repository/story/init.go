package repository

import (
	"stories-backend/internal/domain/story"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type storyRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewStoryRepository(db *mongo.Database, collection *mongo.Collection) domain.StoryRepository {
	return &storyRepository{
		db:         db,
		collection: collection,
	}
}