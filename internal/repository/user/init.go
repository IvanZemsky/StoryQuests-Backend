package repository

import (
	domain "stories-backend/internal/domain/user"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type userRepository struct {
	db         *mongo.Database
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database, collection *mongo.Collection) domain.UserRepository {
	return &userRepository{db: db, collection: collection}
}


