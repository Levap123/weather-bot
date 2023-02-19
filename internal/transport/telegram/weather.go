package telegram

import (
	"Levap123/weather-bot/internal/domain"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gopkg.in/telebot.v3"
)

func (h *Handler) getWeatherByCity(c telebot.Context) error {
	city := c.Message().Payload
	if city == "" {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		var err error
		city, err = h.service.GetCity(ctx, c.Sender().ID)
		if err != nil {
			h.lg.Errorln(err)
			if errors.Is(err, domain.ErrUserNotFound) {
				return c.Send(`мы не нашли вас в нашей базе данных, пожалуйста, поделитесь своей локацией, либо используйте команду с городом.
Например: /weather Astana`)
			}
			return err
		}
	}
	weather, err := h.service.Get(city, viper.GetString("weather_api"))
	if err != nil {
		h.lg.Errorln(err)
		if errors.Is(err, domain.ErrCityNotFound) {
			return c.Send(fmt.Sprintf("Город %s не найден", city))
		}
		return err
	}
	return c.Send(fmt.Sprintf("В городе %s, температура: %0.2f°C, ощущается как: %0.2f°C",
		weather.Location.CityName, weather.Data.Values.Temperature, weather.Data.Values.TemperatureApparent))
}
