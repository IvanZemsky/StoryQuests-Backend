package repository

import (
	domain "stories-backend/internal/domain/story"
	"stories-backend/internal/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (repo *storyRepository) Like(DTO domain.LikeStoryDTO) (domain.LikeStoryResponse, error) {
	ctx, cancel := repository.NewRequestTimeoutContext()
	defer cancel()

	increment := 1
	if DTO.IsLiked {
		increment = -1
	}

	updateOptions := options.FindOneAndUpdate().
		SetReturnDocument(options.After).
		SetUpsert(false)

	var updatedStory domain.Story

	err := repo.collection.FindOneAndUpdate(
		ctx,
		bson.M{"_id": DTO.StoryID},
		bson.M{"$inc": bson.M{"likes": increment}},
		updateOptions,
	).Decode(&updatedStory)

	if err != nil {
		return domain.LikeStoryResponse{}, err
	}

	return domain.LikeStoryResponse{
		StoryID: updatedStory.ID,
		Likes:   updatedStory.Likes,
		IsLiked: !DTO.IsLiked,
	}, nil
}
