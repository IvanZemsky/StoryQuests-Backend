package service

import (
	"errors"
	domain "stories-backend/internal/domain/auth"
	userDomain "stories-backend/internal/domain/user"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

func (s *authService) Register(dto domain.RegisterDTO) (string, error) {
	if err := s.checkIfUserExists(dto.Login); err != nil {
		return "", err
	}

	passwordHash, err := s.hashPassword(dto.Password)
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

	token, err := GererateToken(newUser.ID.String(), dto.Login)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *authService) checkIfUserExists(login string) error {
	_, err := s.userRepo.FindByLogin(login)
	if err == nil {
		return errors.New("user with this login already exists")
	}
	if !errors.Is(err, mongo.ErrNoDocuments) {
		return err
	}

	return nil
}

func (s *authService) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func GererateToken(userID string, login string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := domain.JWTClaims{
		ID:    userID,
		Login: login,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// store secret separate and safe
	tokenString, err := token.SignedString([]byte("secret"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
