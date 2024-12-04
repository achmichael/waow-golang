package cmd

import (
	"fmt"
	"waow-go/internal/routes"
	"waow-go/pkg/database"

	"waow-go/pkg/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func Run() {
	db := database.Connect()
	if db != nil {
		fmt.Println("connected!")
	}
	router := gin.Default()

	router.GET("/", middleware.AuthJWT(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	routes.AuthRoutes(router, db)

	articleRoutes := router.Group("/articles")
	// articleRoutes.Use(middleware.AuthJWT(), middleware.MustUser())
	routes.ArticleRoutes(articleRoutes, db)

	categoryRoutes := router.Group("/categories")
	// categoryRoutes.Use(middleware.AuthJWT(), middleware.MustUser())
	routes.CategoryRoutes(categoryRoutes, db)

	userRoutes := router.Group("/users")
	userRoutes.Use(middleware.AuthJWT(), middleware.MustAdmin())
	routes.UserRoutes(userRoutes, db)
	
	router.Run()
}
