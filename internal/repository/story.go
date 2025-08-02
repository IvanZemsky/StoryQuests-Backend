package repository

import (
	"context"
	"stories-backend/internal/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type storyRepository struct {
	client *mongo.Client
}

func NewStoryRepository(client *mongo.Client) domain.StoryRepository {
	return &storyRepository{client: client}
}

func (repo *storyRepository) Find() ([]domain.Story, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := repo.client.Database("story-quests").Collection("stories")

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
