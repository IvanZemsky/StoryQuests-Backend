package domain

type LoginDTO struct {
	Login    string `bson:"login" json:"login"`
	Password string `bson:"password" json:"password"`
}

type RegisterDTO struct {
	Login    string `bson:"login" json:"login"`
	Password string `bson:"password" json:"password"`
}