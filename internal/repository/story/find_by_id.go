package repository

import (
	"stories-backend/internal/domain/story"
	"stories-backend/internal/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (repo *storyRepository) FindByID(params domain.FindOneStoryParams) (domain.StoryResponse, error) {
	ctx, cancel := repository.NewRequestTimeoutContext()
	defer cancel()

	pipeline := bson.A{
		bson.M{"$match": bson.M{"_id": params.ID}},
	}
	pipeline = append(pipeline, authorPipelineWithoutMatch...)

	if !params.Me.IsZero() {
		pipeline = append(pipeline, getIsLikedPipeline(params.Me)...)
	} else {
		pipeline = append(pipeline, zeroIsLikedPipeline)
	}

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
