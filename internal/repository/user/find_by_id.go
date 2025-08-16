package repository

import (
	domain "stories-backend/internal/domain/user"
	"stories-backend/internal/repository"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func (repo *userRepository) FindByID(id bson.ObjectID) (domain.User, error) {
	ctx, cancel := repository.NewRequestTimeoutContext()
	defer cancel()

	var user domain.User

	err := repo.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}