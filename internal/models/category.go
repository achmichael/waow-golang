package models

import (
	"gorm.io/gorm"
	"waow-go/pkg/common"
)

type Categories struct {
	common.ModelsWithID
	Name        string `json:"name"`
	Description string `json:"description"`
	Slug        string `json:"slug"`

}

func (m *Categories) BeforeCreate(db *gorm.DB) error {
	m.GenerateUUID("category")
	return nil
}

