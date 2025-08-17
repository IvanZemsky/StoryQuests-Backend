package domain

type CreateUserDTO struct {
	Login    string `bson:"login" json:"login"`
	PasswordHash string `bson:"passwordHash" json:"passwordHash"`
}