package routes

import (
	"waow-go/internal/controllers"
	"waow-go/internal/controllers/auth"
	"waow-go/internal/repositories"
	"waow-go/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(route *gin.RouterGroup, db *gorm.DB) {
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)
	authController := auth.NewAuthController(userService)

	route.POST("/", authController.Register)
	route.GET("/", userController.GetUsers)
	route.GET("/:id", userController.GetUserById)
	route.PUT("/:id", userController.UpdateUser)
	route.DELETE("/:id", userController.DeleteUser)
}
