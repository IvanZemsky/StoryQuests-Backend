package service

import (
	domain "stories-backend/internal/domain/auth"
	customErrors "stories-backend/pkg/errors"

	"golang.org/x/crypto/bcrypt"
)

func (s *authService) Login(dto domain.LoginDTO) (string, error) {
	user, err := s.userRepo.FindByLogin(dto.Login)
	if err != nil {
		return "", customErrors.ErrLoginUserNotFound
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(dto.Password)); err != nil {
		if (err == bcrypt.ErrMismatchedHashAndPassword) {
			return "", customErrors.ErrMismatchedPassword
		}
		return "", err
	}

	token, err := generateToken(user.ID.Hex(), user.Login)
	if err != nil {
		return "", err
	}

	return token, nil
}