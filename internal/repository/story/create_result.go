package repository

import (
	"stories-backend/internal/domain/story"
	"stories-backend/internal/repository"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (repo *storyRepository) CreateResult(setResultDTO domain.SetResultDTO) (domain.StoryResult, error) {
	ctx, cancel := repository.NewRequestTimeoutContext()
	defer cancel()

	result := domain.StoryResult{
		ID:       bson.NewObjectID(),
		UserID:   setResultDTO.UserID,
		StoryID:  setResultDTO.StoryID,
		SceneID:  setResultDTO.SceneID,
		Datetime: bson.NewDateTimeFromTime(time.Now()),
	}

	_, err := repo.resultCollection.InsertOne(ctx, result)
	if err != nil {
		return domain.StoryResult{}, err
	}

	return result, nil
}
