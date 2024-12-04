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

type ArticleService struct {
	articleRepository repositories.ArticleInterface
}

func NewArticleService(articleRepository repositories.ArticleInterface) *ArticleService {
	return &ArticleService{articleRepository: articleRepository}
}

func (s *ArticleService) CreateArticle(req *dtos.ArticleRequest) (*common.Message, error) {
	slug := strings.ToLower(req.Title)
	article := &models.Articles{
		Title:       req.Title,
		Content:     req.Content,
		Slug:        slug,
		View_count:  req.View_count,
		User_id:     req.User_id,
		Category_id: req.Category_id,
	}

	err := s.articleRepository.CreateArticle(article)
	if err != nil {
		return nil, err
	}

	response := &common.Message{
		Status:  true,
		Message: "Article created successfully",
	}

	return response, nil
}

func (s *ArticleService) GetAllArticles() (*common.ResponseSuccess, error) {
	article, err := s.articleRepository.GetAllArticles()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("cars tidak ada")
		}

		return nil, err
	}

	resp := &common.ResponseSuccess{
		Status: true,
		Data:   article,
	}

	return resp, nil
}

func (s *ArticleService) GetArticleByID(id string) (*common.ResponseSuccess, error) {
	article, err := s.articleRepository.GetArticleByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("article tidak ada")
		}
		return nil, err
	}

	if article.ID == "" {
		return nil, errors.New("article tidak ada broo")
	}

	resp := &common.ResponseSuccess{
		Status: true,
		Data:   article,
	}

	return resp, nil
}

func (s *ArticleService) UpdateArticle(id string, req *dtos.ArticleRequest) (*common.ResponseSuccess, error) {
	slug := strings.ToLower(req.Title)
	updated_article := &models.Articles{
		Title:       req.Title,
		Content:     req.Content,
		Slug:        slug,
		View_count:  req.View_count,
		User_id:     req.User_id,
		Category_id: req.Category_id,
	}

	article, err := s.articleRepository.GetArticleByID(id)
	if err != nil {
		return nil, err
	}

	if article.ID == "" {
		return nil, errors.New("article tidak ada broo")
	}

	err = s.articleRepository.UpdateArticle(id, updated_article)

	if err != nil {
		return nil, err
	}

	resp := &common.ResponseSuccess{
		Status: true,
		Data:   updated_article,
	}

	return resp, nil
}

func (s *ArticleService) DeleteArticle(id string) (*common.Message, error) {

	article, err := s.articleRepository.GetArticleByID(id)

	if err != nil {
		return nil, err
	}

	if article.ID == "" {
		return nil, errors.New("article tidak ada broo")
	}

	err = s.articleRepository.DeleteArticle(id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("article tidak ada broo")
		}
		return nil, err
	}

	resp := &common.Message{
		Status:  true,
		Message: "Article deleted successfully",
	}

	return resp, nil
}
