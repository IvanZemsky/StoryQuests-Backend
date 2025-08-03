package repository

import (
	"context"
	"stories-backend/internal/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type storyRepository struct {
	db *mongo.Database
}

func NewStoryRepository(db *mongo.Database) domain.StoryRepository {
	return &storyRepository{db: db}
}

func (repo *storyRepository) Find() ([]domain.Story, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := repo.db.Collection("stories")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var stories []domain.Story
	if err = cursor.All(ctx, &stories); err != nil {
		return nil, err
	}

	return stories, nil
}
