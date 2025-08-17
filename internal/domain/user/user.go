package domain

import "go.mongodb.org/mongo-driver/v2/bson"

type User struct {
	ID           bson.ObjectID `bson:"_id" json:"id"`
	Login        string        `bson:"login" json:"login"`
	PasswordHash string        `bson:"passwordHash" json:"passwordHash"`
}

type UserService interface {
	FindByID(id bson.ObjectID) (User, error)
}

type UserRepository interface {
	FindByID(id bson.ObjectID) (User, error)
	FindByLogin(login string) (User, error)
	Create(dto CreateUserDTO) (User, error)
}
