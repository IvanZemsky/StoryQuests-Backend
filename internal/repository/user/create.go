package repository

import (
	domain "stories-backend/internal/domain/user"
	"stories-backend/internal/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (repo *userRepository) Create(dto domain.CreateUserDTO) (domain.User, error) {
	ctx, cancel := repository.NewRequestTimeoutContext()
	defer cancel()

	result, err := repo.collection.InsertOne(ctx, dto)
	if err != nil {
		return domain.User{}, err
	}

	return repo.FindByID(result.InsertedID.(bson.ObjectID))
}