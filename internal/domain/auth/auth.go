package domain

type AuthService interface {
	Login(dto LoginDTO) (string, error)
	Register(dto RegisterDTO) (string, error)
	Logout(token string) error
	GetSession(token string) (Session, error)
}

type Session struct {
	ID string `bson:"_id" json:"id"`
	Login string `bson:"login" json:"login"`
}

