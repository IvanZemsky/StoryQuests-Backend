package service

import (
	domain "stories-backend/internal/domain/auth"

	"golang.org/x/crypto/bcrypt"
)

func (s *authService) Login(dto domain.LoginDTO) (string, error) {
	user, err := s.userRepo.FindByLogin(dto.Login)
	if err != nil {
		return "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(dto.Password)); err != nil {
		// custom error
		return "", err
	}

	token, err := generateToken(user.ID.Hex(), user.Login)
	if err != nil {
		return "", err
	}

	return token, nil
}