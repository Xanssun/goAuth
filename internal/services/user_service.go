package services

import (
	"goauth/config"
	"goauth/internal/models"
	"goauth/pkg/hash"
)

// AuthService предоставляет методы для аутентификации
type AuthService struct{}

// NewAuthService создает новый экземпляр AuthService
func NewAuthService() *AuthService {
	return &AuthService{}
}

// RegisterUser регистрирует нового пользователя
func (s *AuthService) RegisterUser(name, email, password string) (*models.User, error) {
	hash, err := hash.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Name:     name,
		Email:    email,
		Password: hash,
	}

	result := config.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
