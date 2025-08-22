package service

import (
	domain "stories-backend/internal/domain/story"
)

// better separate adding and removing likes in both here and in repository
func (service *storyService) Like(DTO domain.LikeStoryDTO) (domain.LikeStoryResponse, error) {
	_, err := service.repo.FindByID(DTO.StoryID)
	if err != nil {
		return domain.LikeStoryResponse{}, err
	}
	
	if !DTO.IsLiked {
		err := service.likeRepo.AddLike(DTO.StoryID, DTO.UserID)
		if err != nil {
			return domain.LikeStoryResponse{}, err
		}
	}

	if DTO.IsLiked {
		err := service.likeRepo.RemoveLike(DTO.StoryID, DTO.UserID)
		if err != nil {
			return domain.LikeStoryResponse{}, err
		}
	}
	
	res, err := service.repo.Like(DTO)
	if err != nil {
		return domain.LikeStoryResponse{}, err
	}
	
	return res, nil
}