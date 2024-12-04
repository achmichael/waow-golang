package services

import (
	"errors"
	"fmt"
	"strings"
	"waow-go/internal/dtos"
	"waow-go/internal/models"
	"waow-go/internal/repositories"
	"waow-go/pkg/common"

	"gorm.io/gorm"
)

type CategoryService struct {
	categoryRepository repositories.CategoryInterface
}

func NewCategoryService(categoryRepository repositories.CategoryInterface) *CategoryService {
	return &CategoryService{categoryRepository: categoryRepository}
}

func (s *CategoryService) CreateCategory(req *dtos.CategoryRequest) (*common.ResponseSuccess, error) {
	slug := strings.ToLower(req.Name)
	category := &models.Categories{
		Name:        req.Name,
		Description: req.Description,
		Slug:        slug,
	}

	err := s.categoryRepository.CreateCategory(category)

	if err != nil {
		return nil, err
	}

	response := &common.ResponseSuccess{
		Status: true,
		Data:   category,
	}

	return response, nil
}

func (s *CategoryService) GetAllCategories() (*common.ResponseSuccess, error) {
	categories, err := s.categoryRepository.GetAllCategories()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("category tidak ada")
		}

		return nil, err
	}

	response := &common.ResponseSuccess{
		Status: true,
		Data:   categories,
	}

	return response, nil
}

func (s *CategoryService) GetCategoryByID(id string) (*common.ResponseSuccess, error) {
	category, err := s.categoryRepository.GetCategoryByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("category tidak ada")
		}
		return nil, err
	}

	if category == nil {
		return nil, errors.New("category tidak ada")
	}
	if category.ID == "" {
		return nil, fmt.Errorf("category tidak ada")
	}

	response := &common.ResponseSuccess{
		Status: true,
		Data:   category,
	}

	return response, nil
}

func (s *CategoryService) UpdateCategory(id string, req *dtos.CategoryRequest) (*common.ResponseSuccess, error) {
	slug := strings.ToLower(req.Name)
	updated_category := &models.Categories{
		Name:        req.Name,
		Description: req.Description,
		Slug:        slug,
	}

	category, err := s.categoryRepository.GetCategoryByID(id)

	if err != nil {
		return nil, err
	}

	if category.ID == "" {
		return nil, errors.New("category tidak ada")
	}

	err = s.categoryRepository.UpdateCategory(id, updated_category)

	if err != nil {
		return nil, err
	}

	response := &common.ResponseSuccess{
		Status: true,
		Data:   category,
	}
	return response, nil
}

func (s *CategoryService) DeleteCategory(id string) (*common.Message, error) {

	category, err := s.categoryRepository.GetCategoryByID(id)

	if err != nil {
		return nil, err
	}

	if category.ID == "" {
		return nil, errors.New("category tidak ada")
	}

	err = s.categoryRepository.DeleteCategory(id)

	if err != nil {
		return nil, err
	}

	response := &common.Message{
		Status:  true,
		Message: "category berhasil dihapus",
	}

	return response, nil
}
