package service

import (
	"errors"
	domain "stories-backend/internal/domain/auth"
	userDomain "stories-backend/internal/domain/user"
	commonErrors "stories-backend/pkg/errors"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

func (s *authService) Register(dto domain.RegisterDTO) (string, error) {
	if err := s.checkIfUserExists(dto.Login); err != nil {
		return "", err
	}

	passwordHash, err := hashPassword(dto.Password)
	if err != nil {
		return "", err
	}

	newUser, err := s.userRepo.Create(userDomain.CreateUserDTO{
		Login:        dto.Login,
		PasswordHash: passwordHash,
	})
	if err != nil {
		return "", err
	}

	token, err := generateToken(newUser.ID.Hex(), dto.Login)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *authService) checkIfUserExists(login string) error {
	_, err := s.userRepo.FindByLogin(login)
	if err == nil {
		return commonErrors.ErrUserAlreadyExists
	}
	if !errors.Is(err, mongo.ErrNoDocuments) {
		return err
	}

	return nil
}
