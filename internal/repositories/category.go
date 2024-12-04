package repositories

import (
	"fmt"
	"waow-go/internal/models"

	"gorm.io/gorm"
)

type CategoryInterface interface {
	CreateCategory(*models.Categories) error
	GetAllCategories() ([]*models.Categories, error)
	GetCategoryByID(id string) (*models.Categories, error)
	UpdateCategory(id string, category *models.Categories) error
	DeleteCategory(id string) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) CreateCategory(category *models.Categories) error {
	return r.db.Create(category).Error
}


func (r *categoryRepository) GetAllCategories() ([]*models.Categories, error) {
	categories := []*models.Categories{}
	err := r.db.Find(&categories).Error

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *categoryRepository) GetCategoryByID(id string) (*models.Categories, error){
	category := &models.Categories{}
	err := r.db.Find(&category, "id = ?", id).Error
	sql := r.db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Find(&category, "id = ?", id)
	})

	fmt.Println(sql)
	return category, err
}

func (r *categoryRepository) UpdateCategory(id string, category *models.Categories) error {
	return r.db.Where("id = ?", id).Updates(category).Find(&category, "id = ?", id).Error
}

func (r *categoryRepository) DeleteCategory(id string) error {
	// return r.db.Where("id = ?", id).Delete(&models.Categories{}).Error
	tx := r.db.Begin()
    defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Where("id = ?", id).Delete(&models.Categories{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}
	
	return nil
}
