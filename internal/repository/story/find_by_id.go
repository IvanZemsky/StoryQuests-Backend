package repository

import (
	"stories-backend/internal/domain/story"
	"stories-backend/internal/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (repo *storyRepository) FindByID(id bson.ObjectID) (domain.StoryResponse, error) {
	ctx, cancel := repository.NewRequestTimeoutContext()
	defer cancel()

	pipeline := bson.A{
		bson.M{"$match": bson.M{"_id": id}},
	}
	pipeline = append(pipeline, authorPipelineWithoutMatch...)

	cursor, err := repo.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return domain.StoryResponse{}, err
	}
	defer cursor.Close(ctx)

	if cursor.Next(ctx) {
		var response domain.StoryResponse
		if err := cursor.Decode(&response); err != nil {
			return domain.StoryResponse{}, err
		}
		return response, nil
	}

	return domain.StoryResponse{}, mongo.ErrNoDocuments
}
