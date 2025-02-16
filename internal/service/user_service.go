package service

import (
	"context"
	"github.com/codeboris/avito-shop/internal/models"
	"github.com/codeboris/avito-shop/internal/repository"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) AuthenticateUser(ctx context.Context, username string) (*models.User, error) {
	user, err := s.userRepo.GetOrCreateUser(ctx, username)
	if err != nil {
		return nil, err
	}

	return user, nil
}
