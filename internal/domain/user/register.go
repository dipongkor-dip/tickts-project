package user

import (
	"gotickts/internal/auth"
	"gotickts/internal/config"
	"gotickts/internal/middlewares"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB, cfg *config.Config) {
	jwtService := auth.NewJWTService(cfg.JWT_secret)

	userRepository := NewRepository(db)
	userService := NewService(userRepository, jwtService)
	userHandler := NewHandler(userService)

	api := e.Group("/api/v1/auth")

	api.POST("/register", userHandler.CreateUser)
	api.POST("/login", userHandler.LoginUser)
	api.GET("/me", userHandler.GetMe, middlewares.AuthMiddleware(jwtService)) // protected route
	api.GET("", userHandler.getUsers)
}
