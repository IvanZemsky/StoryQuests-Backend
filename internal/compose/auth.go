package compose

import (
	userDomain "stories-backend/internal/domain/user"
	authDomain "stories-backend/internal/domain/auth"
	"stories-backend/internal/handlers/auth"
	"stories-backend/internal/service/auth"
)

func InitAuthModule(
	init InitModuleOptions,
	userRepo userDomain.UserRepository,
	jwt authDomain.JWTConfig,
) {
	authService := service.NewAuthService(userRepo, jwt)
	handlers.NewAuthHandler(init.Router, authService)
}
