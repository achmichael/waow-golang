package controllers

import (
	"waow-go/internal/dtos"
	"waow-go/internal/services"
	"waow-go/pkg/common"

	"github.com/gin-gonic/gin"
)

type categoryController struct {
	categoryService *services.CategoryService
}

func NewCategoryController(categoryService *services.CategoryService) *categoryController {
	return &categoryController{categoryService: categoryService}
}

func (c *categoryController) CreateCategory(ctx *gin.Context) {
	req := new(dtos.CategoryRequest)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		response := &common.ResponseError{
			Status:  false,
			Message: err.Error(),
		}
		ctx.JSON(400, response)
		return
	}

	resp, err := c.categoryService.CreateCategory(req)
	if err != nil {
		response := &common.ResponseError{		
			Status:  false,
			Message: err.Error(),
		}
		ctx.JSON(400, response)
		return
	}

	ctx.JSON(201, resp)
}

func (c *categoryController) GetAllCategories(ctx *gin.Context) {
	resp, err := c.categoryService.GetAllCategories()
	if err != nil {
		response := &common.ResponseError{
			Status:  false,
			Message: err.Error(),
		}
		ctx.JSON(400, response)
		return
	}
	ctx.JSON(200, resp)
}

func (c *categoryController) GetCategoryByID(ctx *gin.Context) {
	id := ctx.Param("id")
	resp, err := c.categoryService.GetCategoryByID(id)
	if err != nil {
		response := &common.ResponseError{
			Status:  false,
			Message: err.Error(),
		}
		ctx.JSON(400, response)
		return
	}
	ctx.JSON(200, resp)
}

func (c *categoryController) UpdateCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	req := new(dtos.CategoryRequest)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		response := &common.ResponseError{
			Status:  false,
			Message: err.Error(),
		}
		ctx.JSON(400, response)
		return
	}

	resp, err := c.categoryService.UpdateCategory(id, req)

	if err != nil {
		response := &common.ResponseError{
			Status:  false,
			Message: err.Error(),
		}
		ctx.JSON(400, response)
		return
	}

	ctx.JSON(200, resp)
}

func (c *categoryController) DeleteCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	resp, err := c.categoryService.DeleteCategory(id)
	if err != nil {
		response := &common.ResponseError{
			Status:  false,
			Message: err.Error(),
		}
		ctx.JSON(400, response)
		return
	}

	ctx.JSON(200, resp)
}
