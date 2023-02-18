package telegram

import (
	"Levap123/weather-bot/internal/entity"
	"Levap123/weather-bot/internal/service"
	"context"
)

type Service struct {
	WeatherService
	UserService
}

type WeatherService interface {
	Get(city, weatherApi string) (*entity.Weather, error)
}

type UserService interface {
	Create(ctx context.Context, userID int64, lng, lat float32) error
	GetCity(ctx context.Context, userID int64) (string, error)
}

func NewService(strg *service.Storage) *Service {
	return &Service{
		WeatherService: service.NewWeatherService(),
		UserService:    service.NewUserService(strg.UserRepo),
	}
}
