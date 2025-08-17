package domain



type AuthService interface {
	Login(dto LoginDTO) (string, error)
	Register(dto RegisterDTO) (string, error)
	Logout(token string) error
	GetSession(token string) (string, error)
}

