package controllers

import (
	"net/http"
	"waow-go/internal/dtos"
	"waow-go/internal/services"
	"github.com/gin-gonic/gin"
)

type articleController struct {
	articleService *services.ArticleService
}

func NewArticleController(articleService *services.ArticleService) *articleController {
   return &articleController{articleService : articleService }
}

func (c *articleController) CreateArticle(ctx *gin.Context) {
	req := new(dtos.ArticleRequest)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := c.articleService.CreateArticle(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, resp)
}

func (c *articleController) GetAllArticles(ctx *gin.Context) {
	resp , err := c.articleService.GetAllArticles()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (c *articleController) GetArticleByID(ctx *gin.Context) {
	id := ctx.Param("id")
	resp, err := c.articleService.GetArticleByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (c *articleController) UpdateArticle (ctx *gin.Context) {
	id := ctx.Param("id")
	req := new(dtos.ArticleRequest)
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := c.articleService.UpdateArticle(id, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (c *articleController) DeleteArticle (ctx *gin.Context) {
	id := ctx.Param("id")
	resp, err := c.articleService.DeleteArticle(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, resp)
}