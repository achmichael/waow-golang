package routes

import (
	"waow-go/internal/controllers"
	"waow-go/internal/repositories"
	"waow-go/internal/services"
	"waow-go/pkg/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func CategoryRoutes(route *gin.RouterGroup, db *gorm.DB) {
	categoryRepository := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepository)
	categoryController := controllers.NewCategoryController(categoryService)

	route.POST("/", middleware.AuthJWT(), middleware.MustUser(),categoryController.CreateCategory)
	route.GET("/", categoryController.GetAllCategories)
	route.GET("/:id", categoryController.GetCategoryByID)
	route.PUT("/:id", categoryController.UpdateCategory)
	route.DELETE("/:id", categoryController.DeleteCategory)
}