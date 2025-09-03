package service

import (
	"errors"
	"log"
	domain "stories-backend/internal/domain/story"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (s *storyService) SetResult(setResultDTO domain.SetResultDTO) (domain.StoryResult, error) {
	existingResult, err := s.repo.FindResultByUserIDAndStoryID(setResultDTO.UserID, setResultDTO.StoryID)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return s.repo.CreateResult(setResultDTO)
		}

		// Любая другая ошибка - логируем и возвращаем
		log.Println("FindResultByUserIDAndStoryID err", err)
		return domain.StoryResult{}, err
	}

	log.Println("existingResult found with ID:", existingResult.ID)

	// Если документ найден - обновляем его
	return s.repo.UpdateResult(setResultDTO)
}
