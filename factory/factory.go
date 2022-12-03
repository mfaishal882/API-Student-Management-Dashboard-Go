package factory

import (
	authDelivery "api-alta-dashboard/features/auth/delivery"
	authRepo "api-alta-dashboard/features/auth/repository"
	authService "api-alta-dashboard/features/auth/service"

	userDelivery "api-alta-dashboard/features/user/delivery"
	userRepo "api-alta-dashboard/features/user/repository"
	userService "api-alta-dashboard/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	userRepoFactory := userRepo.New(db)
	// userRepoFactory := userRepo.NewRaw(db)
	userServiceFactory := userService.New(userRepoFactory)
	userDelivery.New(userServiceFactory, e)

	authRepoFactory := authRepo.New(db)
	authServiceFactory := authService.New(authRepoFactory)
	authDelivery.New(authServiceFactory, e)

}