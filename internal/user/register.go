package user

import (
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	userRepository := NewRepository(db)
	userService := NewService(userRepository)
	userHandler := NewHandler(userService)
	api := e.Group("/api/v1/auth")
	// APi endpoint
	api.POST("/register", userHandler.CreateUser) //api/vi/auth/register
	api.POST("/login", userHandler.LoginUser)    //api/vi/auth/login
}

