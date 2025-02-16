package service

import (
	"errors"
	"github.com/codeboris/avito-shop/internal/models"
	"github.com/codeboris/avito-shop/internal/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// Register регистрирует нового пользователя
func (s *UserService) Register(username, password string) (*models.User, error) {
	user := &models.User{
		Username: username,
		Password: password,
		Coins:    0, // Начальное количество монет
	}

	if err := s.userRepo.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

// Authenticate аутентифицирует пользователя
func (s *UserService) Authenticate(username, password string) (*models.User, error) {
	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	if user.Password != password {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

// GetUserInfo возвращает информацию о пользователе
func (s *UserService) GetUserInfo(userID int) (*models.User, error) {
	return s.userRepo.GetUserByID(userID)
}

// UpdateUserCoins обновляет количество монет пользователя
func (s *UserService) UpdateUserCoins(userID, coins int) error {
	return s.userRepo.UpdateUserCoins(userID, coins)
}
