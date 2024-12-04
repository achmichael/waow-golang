package models

import (
	"gorm.io/gorm"
	"waow-go/pkg/common"
)

type Articles struct {
	common.ModelsWithID
	Category_id string `json:"category_id" `
	User_id     string `json:"user_id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Slug        string `json:"slug"`
	View_count  int    `json:"view_count"`

	Category Categories `json:"category" gorm:"foreignKey:Category_id;references:ID"`
	User     Users      `json:"user" gorm:"foreignKey:User_id;references:ID"`
}

func (m *Articles) BeforeCreate(db *gorm.DB) error {
	m.GenerateUUID("article")
	return nil
}
