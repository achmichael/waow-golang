package services

import (
	"errors"
	"time"
	"waow-go/internal/dtos"
	"waow-go/internal/models"
	"waow-go/internal/repositories"
	"waow-go/pkg/auth"
	"waow-go/pkg/common"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	userRepository repositories.UserInterface
}

func NewUserService(repo repositories.UserInterface) *UserService {
	return &UserService{userRepository: repo}
}

func (s *UserService) Login(req dtos.LoginRequest) (*common.ResponseSuccess, error) {
	user, err := s.userRepository.Login(req.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("password is incorrect")
	}

	authToken, err := auth.GenerateJWT(user.Username, user.Role)
	
	if err != nil {
		return nil, err
	}

	response := &common.ResponseSuccess{
		Status: true,
		Data:   authToken,
	}

	return response, nil
}

func (s *UserService) GetAllUsers() (*common.ResponseSuccess, error) {
	users, err := s.userRepository.GetAllUsers()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("users not found")
		}
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New("data users tidak ada")
	}

	response := &common.ResponseSuccess{
		Status: true,
		Data:   users,
	}

	return response, nil
}

func (s *UserService) Register(req dtos.RegisterRequest) (*common.Message, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	
	if err != nil {
		return nil, err
	}

	err = s.userRepository.Register(&models.Users{
		Username:            req.Username,
		Password:            string(hashedPassword),
		Email:               req.Email,
		DisplayName:         req.DisplayName,
		Bio:                 req.Bio,
		Profile_Picture_Url: "#",
		Role:                req.Role,
		Registration_date:   time.Now(),
	})

	if err != nil {
		return nil, errors.New("failed to create user")
	}

	response := &common.Message{
		Status:  true,
		Message: "User created successfully",
	}

	return response, nil
}


func (s *UserService) GetUserById(id string) (*common.ResponseSuccess, error) {
	user, err := s.userRepository.GetUserById(id)
	
	if err != nil {
		return nil, err
	}
	
	if user.ID == "" {
		return nil, errors.New("user not found bro")
	}

	response := &common.ResponseSuccess {
		Status: true,
		Data: user,
	}

	return response, nil
}

func (s *UserService) UpdateUser(id string, req *dtos.RegisterRequest) (*common.ResponseSuccess, error) {
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	updated_user := &models.Users{
		Username: req.Username,
		Password: string(hashed_password),
		Email: req.Email,
		DisplayName: req.DisplayName,
		Bio: req.Bio,
		Profile_Picture_Url: "#",
	}

	user, err := s.userRepository.GetUserById(id)
	
	if err != nil {
		return nil, err
	}

	if user.ID == "" {
		return nil, errors.New("user not found")
	}

	err = s.userRepository.Update(id, updated_user)

	if err != nil {
		return nil, err
	}

	response := &common.ResponseSuccess{
		Status: true,
		Data: map[string]interface{}{
			"user_id": user.ID,
			"message": "User updated successfully",
		},
	}
	return response, nil
}

func (s *UserService) DeleteUser(id string) (*common.Message, error) {
	user, err := s.userRepository.GetUserById(id)
	if err != nil {
		return nil, err
	}

	if user.ID == "" {
		return nil, errors.New("user not found")
	}

	err = s.userRepository.Delete(id)

	if err != nil {
		return nil, err
	}

	response := &common.Message{
		Status:  true,
		Message: "User deleted successfully",
	}

	return response, nil
}