package repository

import (
	"stories-backend/internal/domain/scene"
	"stories-backend/internal/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (repo *sceneRepository) CreateForStory(storyID bson.ObjectID, dto []domain.CreateSceneDTO) error {
	ctx, cancel := repository.NewRequestTimeoutContext()
	defer cancel()

	scenes := make([]domain.Scene, len(dto))
	for i, v := range dto {
		scenes[i] = createFromDTO(&v, storyID)
	}

	_, err := repo.collection.InsertMany(ctx, scenes)
	if err != nil {
		return err
	}

	return nil
}

func createFromDTO(dto *domain.CreateSceneDTO, storyID bson.ObjectID) domain.Scene {
	return domain.Scene{
		Title:       dto.Title,
		Description: dto.Description,
		Img:         dto.Img,
		Type:        dto.Type,
		Answers:     dto.Answers,
		StoryID:     storyID,
		Number:      dto.Number,
	}
}
