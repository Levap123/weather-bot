package service

import (
	"Levap123/weather-bot/internal/entity"
	"context"
)

type UserService struct {
	repo UserRepo
}

func NewUserService(repo UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) Create(ctx context.Context, user *entity.User) error {
	return us.repo.Create(ctx, user)
}

func (us *UserService) GetCity(ctx context.Context, userID string) (string, error) {
	return us.repo.GetCity(ctx, userID)
}
