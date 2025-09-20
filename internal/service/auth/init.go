package service

import (
	authDomain "stories-backend/internal/domain/auth"
	userDomain "stories-backend/internal/domain/user"
)

type authService struct {
	userRepo userDomain.UserRepository
	jwt authDomain.JWTConfig
}

func NewAuthService(userRepo userDomain.UserRepository, jwt authDomain.JWTConfig) authDomain.AuthService {
	return &authService{userRepo: userRepo, jwt: jwt}
}
