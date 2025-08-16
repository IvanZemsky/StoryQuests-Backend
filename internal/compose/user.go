package compose

import (
	"stories-backend/internal/domain/user"
	"stories-backend/internal/handlers/user"
	"stories-backend/internal/service/user"
)


func InitUserModule(
	init InitModuleOptions,
	userRepo domain.UserRepository,
) {
	userService := service.NewUserService(userRepo)
	handlers.NewUserHandler(init.Router, userService)
}
