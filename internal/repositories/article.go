package repositories

import (
	"waow-go/internal/models"

	"gorm.io/gorm"
)

type ArticleInterface interface {
	CreateArticle(*models.Articles) error
	GetAllArticles() ([]*models.Articles, error)
	GetArticleByID(id string) (*models.Articles, error)
	UpdateArticle(id string, article *models.Articles) error
	DeleteArticle(id string) error
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *articleRepository {
	return &articleRepository{db: db}
}

func (r *articleRepository) CreateArticle(article *models.Articles) error {
	return r.db.Create(article).Error
}

func (r *articleRepository) GetAllArticles() ([]*models.Articles, error) {
	articles := []*models.Articles{}
	err := r.db.Find(&articles).Error
	return articles, err
}

func (r *articleRepository) GetArticleByID(id string) (*models.Articles, error) {
	article := &models.Articles{}
	err := r.db.Find(article, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return article, err
}

func (r *articleRepository) UpdateArticle(id string, article *models.Articles) error {
	return r.db.Where("id = ?", id).Updates(article).Find(&article, "id = ?", id).Error
}

func (r *articleRepository) DeleteArticle(id string) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Where("id = ?", id).Delete(&models.Articles{}).Error
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
