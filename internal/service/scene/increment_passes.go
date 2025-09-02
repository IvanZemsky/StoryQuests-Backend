package service

import (
	"errors"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (s *sceneService) IncrementPasses(sceneID bson.ObjectID) error {
	scene, err := s.repo.FindByID(sceneID)
	if err != nil {
		return err
	}
	
	if scene.Type != "end" {
		return errors.New("scene is not an end scene")
	}

	return s.repo.IncrementPasses(sceneID)
}