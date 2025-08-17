package service

import (
	authDomain "stories-backend/internal/domain/auth"
	userDomain "stories-backend/internal/domain/user"
)

type authService struct {
	userRepo userDomain.UserRepository
}

func NewAuthService(userRepo userDomain.UserRepository) authDomain.AuthService {
	return &authService{userRepo: userRepo}
}
