package repositories

import (
	"waow-go/internal/models"
	"errors"
	"gorm.io/gorm"
)

type UserInterface interface {
	Register(*models.Users) error
	Login(username string) (*models.Users, error)
	GetAllUsers() ([]*models.Users, error)
	GetUserById(id string) (*models.Users, error)
	Update(id string, user *models.Users) error
	Delete(id string) error
}

type userRepository struct {
	dB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{dB: db}
}

func (r *userRepository) GetAllUsers() ([]*models.Users, error) {
	users := []*models.Users{}

	err := r.dB.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) Register(user *models.Users) error {
	return r.dB.Create(user).Error
}

func (r *userRepository) Login(username string) (*models.Users, error) {
	user := &models.Users{}

	err := r.dB.Find(user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) GetUserById(id string) (*models.Users, error) {
	user := &models.Users{}

	err := r.dB.Find(user, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return user, nil
}


func (r *userRepository) Update(id string, user *models.Users) error {
	err := r.dB.Where("id = ?", id).Updates(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Delete(id string) error {
	err := r.dB.Where("id = ?", id).Delete(&models.Users{}).Error
	if err != nil {
		return err
	}

	return nil
}

