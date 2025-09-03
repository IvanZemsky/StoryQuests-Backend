package repository

import (
	"stories-backend/internal/domain/story"
	"stories-backend/internal/repository"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (repo *storyRepository) UpdateResult(setResultDTO domain.SetResultDTO) (domain.StoryResult, error) {
	ctx, cancel := repository.NewRequestTimeoutContext()
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"sceneId":  setResultDTO.SceneID,
			"datetime": bson.NewDateTimeFromTime(time.Now()),
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var result domain.StoryResult
	err := repo.resultCollection.FindOneAndUpdate(
		ctx,
		bson.M{
			"userId":  setResultDTO.UserID,
			"storyId": setResultDTO.StoryID,
		},
		update,
		opts,
	).Decode(&result)

	if err != nil {
		return domain.StoryResult{}, err
	}

	return result, nil
}
