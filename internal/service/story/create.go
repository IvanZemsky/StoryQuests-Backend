package service

import (
	sceneDomain "stories-backend/internal/domain/scene"
	"stories-backend/internal/domain/story"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (s *storyService) Create(
	storyDTO domain.CreateStoryDTO,
	scenesDTO []sceneDomain.CreateSceneDTO,
) (bson.ObjectID, error) {
	storyID, err := s.repo.Create(storyDTO)
	if err != nil {
		return bson.ObjectID{}, err
	}

	err = s.sceneRepo.CreateForStory(storyID, scenesDTO)
	if err != nil {
		return bson.ObjectID{}, err
	}

	return storyID, nil
}
