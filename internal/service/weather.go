package service

import (
	"Levap123/weather-bot/internal/domain"
	"Levap123/weather-bot/internal/entity"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type WeatherService struct{}

func NewWeatherService() *WeatherService {
	return &WeatherService{}
}

func (w *WeatherService) Get(city, weatherApi string) (*entity.Weather, error) {
	url := fmt.Sprintf(`https://api.tomorrow.io/v4/weather/realtime?location=%s&apikey=%s`, city, weatherApi)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("service - get weather %w", err)
	}

	req.Header.Add("accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("service - get weather %w", err)
	}

	if res.StatusCode == http.StatusBadRequest {
		return nil, fmt.Errorf("service - get weather %w", domain.ErrCityNotFound)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("service - get weather %w", err)
	}
	fmt.Println(string(body))
	var weather entity.Weather
	if err := json.Unmarshal(body, &weather); err != nil {
		return nil, fmt.Errorf("service - get weather %w", err)
	}

	return &weather, nil
}
