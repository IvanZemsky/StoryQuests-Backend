package repository

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
	"stories-backend/internal/domain/story"
)

type storyRepository struct {
	db               *mongo.Database
	collection       *mongo.Collection
	resultCollection *mongo.Collection
}

func NewStoryRepository(db *mongo.Database,
	collection *mongo.Collection,
	resultCollection *mongo.Collection,
) domain.StoryRepository {
	return &storyRepository{
		db:               db,
		collection:       collection,
		resultCollection: resultCollection,
	}
}
