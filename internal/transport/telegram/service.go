package telegram

import (
	"Levap123/weather-bot/internal/entity"
	"Levap123/weather-bot/internal/service"
)

type Service struct {
	WeatherService
}

type WeatherService interface {
	Get(city, weatherApi string) (*entity.Weather, error)
}

func NewService() *Service {
	return &Service{
		WeatherService: service.NewWeatherService(),
	}
}
