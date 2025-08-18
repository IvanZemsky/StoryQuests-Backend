package service

import domain "stories-backend/internal/domain/auth"

func (s *authService) GetSession(token string) (domain.Session, error) {
	session, err := ValidateToken(token)
	if err != nil {
		return domain.Session{}, err
	}

	return domain.Session{ID: session.ID, Login: session.Login}, nil
}
