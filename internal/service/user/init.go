package service

import (
	domain "stories-backend/internal/domain/user"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type userService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) domain.UserService {
	return &userService{repo: repo}
}

func (s *userService) FindByID(id bson.ObjectID) (domain.User, error) {
	return s.repo.FindByID(id)
}
