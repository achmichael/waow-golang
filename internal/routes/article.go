package routes

import (
	"waow-go/internal/controllers"
	"waow-go/internal/repositories"
	"waow-go/internal/services"
	"waow-go/pkg/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ArticleRoutes(route *gin.RouterGroup, db *gorm.DB) {
	article := repositories.NewArticleRepository(db)
	articleService := services.NewArticleService(article)
	articleController := controllers.NewArticleController(articleService)

	route.POST("/", middleware.AuthJWT(), middleware.MustUser() ,articleController.CreateArticle)
	route.GET("/", articleController.GetAllArticles)
	route.GET("/:id", articleController.GetArticleByID)
	route.PUT("/:id", articleController.UpdateArticle)
	route.DELETE("/:id", articleController.DeleteArticle)
}
