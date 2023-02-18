package service

import (
	"Levap123/weather-bot/internal/entity"
	"Levap123/weather-bot/internal/repository/mongor"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Storage struct {
	UserRepo
}

type UserRepo interface {
	Create(ctx context.Context, user *entity.User) error
	GetCity(ctx context.Context, userID int64) (string, error)
}

func NewStorageMongo(cl *mongo.Client) *Storage {
	return &Storage{
		UserRepo: mongor.NewUser(cl),
	}
}
