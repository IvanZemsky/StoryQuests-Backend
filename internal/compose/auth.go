package compose

import (
	userDomain "stories-backend/internal/domain/user"
	"stories-backend/internal/handlers/auth"
	"stories-backend/internal/service/auth"
)

func InitAuthModule(
	init InitModuleOptions,
	userRepo userDomain.UserRepository,
) {
	authService := service.NewAuthService(userRepo)
	handlers.NewAuthHandler(init.Router, authService)
}
