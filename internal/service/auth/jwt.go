package service

import (
	"errors"
	domain "stories-backend/internal/domain/auth"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func generateToken(userID string, login string) (string, error) {
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

func ValidateToken(tokenStr string) (domain.JWTClaims, error) {
    claims := domain.JWTClaims{}
    token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (any, error) {
        return []byte("secret"), nil
    })

    if err != nil {
        return domain.JWTClaims{}, err
    }

    if claims, ok := token.Claims.(*domain.JWTClaims); ok && token.Valid {
        return *claims, nil
    } else {
        return domain.JWTClaims{}, errors.New("invalid token")
    }
}

// func validateToken(tokenStr string) (domain.JWTClaims, error) {
// 	claims := &domain.JWTClaims{}
// 	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
// 		return []byte("secret"), nil
// 	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

// 	if err != nil {
// 		return domain.JWTClaims{}, err
// 	}

// 	if claims, ok := token.Claims.(domain.JWTClaims); ok {
// 		return claims, nil
// 	} else {
// 		return domain.JWTClaims{}, errors.New("invalid token")
// 	}

// }
