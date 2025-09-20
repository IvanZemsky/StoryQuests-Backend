package repository

import (
	"stories-backend/internal/domain/story"
	"stories-backend/internal/repository"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (repo *storyRepository) Create(DTO *domain.CreateStoryDTO) (bson.ObjectID, error) {
	ctx, cancel := repository.NewRequestTimeoutContext()
	defer cancel()

	story := createFromDTO(DTO)

	result, err := repo.collection.InsertOne(ctx, story)
	if err != nil {
		return bson.ObjectID{}, err
	}

	return result.InsertedID.(bson.ObjectID), nil
}

func createFromDTO(dto *domain.CreateStoryDTO) domain.Story {
	return domain.Story{
		Name:        dto.Name,
		Description: dto.Description,
		AuthorID:    dto.AuthorID,
		SceneCount:  dto.SceneCount,
		Img:         dto.Img,
		Tags:        dto.Tags,
		Date:        time.Now(),
		Likes:       0,
		Passes:      0,
	}
}
