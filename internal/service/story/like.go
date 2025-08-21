package service

import (
	domain "stories-backend/internal/domain/story"
)

func (service *storyService) Like(DTO domain.LikeStoryDTO) (domain.LikeStoryResponse, error) {
	res, err := service.repo.Like(DTO)
	if err != nil {
		return domain.LikeStoryResponse{}, err
	}
	return res, nil
}