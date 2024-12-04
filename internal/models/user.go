package models

import (
	"time"
	"waow-go/pkg/common"

	"gorm.io/gorm"
)

const (
	ROLE_ADMIN = "ADMIN"
	ROLE_USER  = "USER"
)

type Users struct {
	common.ModelsWithID           
	Username            string `json:"username" gorm:"unique"`
	Password            string `json:"password"`
	Email               string `json:"email" gorm:"unique"`
	Role                string `json:"role"`
	DisplayName         string `json:"displayName"`
	Bio                 string `json:"bio"`
	Profile_Picture_Url string `json:"profile_picture_url" gorm:"default:null"`
	Registration_date   time.Time `json:"registration_date"`
}

func (m *Users) BeforeCreate(db * gorm.DB) error {
	m.GenerateUUID("user")
	return nil
}
