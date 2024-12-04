package routes

import (
	"waow-go/internal/controllers/auth"
	"waow-go/internal/repositories"
	"waow-go/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRoutes(route *gin.Engine, db *gorm.DB) {
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	loginController := auth.NewAuthController(userService)

	route.POST("/login", loginController.Login)
	route.POST("/register", loginController.Register)
}
