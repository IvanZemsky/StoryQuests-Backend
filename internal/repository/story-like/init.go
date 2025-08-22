package repository

import (
	domain "stories-backend/internal/domain/story"
	"stories-backend/internal/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type storyLikeRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewStoryLikeRepository(db *mongo.Database, collection *mongo.Collection) domain.StoryLikeRepository {
	return &storyLikeRepository{
		db:         db,
		collection: collection,
	}
}

func (repo *storyLikeRepository) AddLike(storyID bson.ObjectID, userID bson.ObjectID) error {
	ctx, cancel := repository.NewRequestTimeoutContext()
	defer cancel()

	_, err := repo.collection.InsertOne(ctx, domain.StoryLike{
		StoryID: storyID,
		UserID:  userID,
	})
	return err
}

func (repo *storyLikeRepository) RemoveLike(storyID bson.ObjectID, userID bson.ObjectID) error {
	ctx, cancel := repository.NewRequestTimeoutContext()
	defer cancel()

	_, err := repo.collection.DeleteOne(ctx, bson.M{"storyId": storyID, "userId": userID})
	return err
}

func (repo *storyLikeRepository) FindLikes(
	storyID bson.ObjectID,
	userID bson.ObjectID,
) ([]domain.LikeStoryResponse, error) {
	return nil, nil
}

func (repo *storyLikeRepository) FindLike(
	storyID bson.ObjectID,
	userID bson.ObjectID,
) ([]domain.LikeStoryResponse, error) {
	return nil, nil
}
