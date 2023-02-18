package service

import (
	"Levap123/weather-bot/internal/entity"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserService struct {
	repo UserRepo
}

func NewUserService(repo UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) Create(ctx context.Context, userID int64, lng, lat float32) error {
	user := &entity.User{
		ID: userID,
	}
	query := fmt.Sprintf(`https://api.bigdatacloud.net/data/reverse-geocode-client?latitude=%f&longitude=%f&localityLanguage=en`, lat, lng)
	req, err := http.NewRequest("GET", query, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return err
	}

	return us.repo.Create(ctx, user)
}

func (us *UserService) GetCity(ctx context.Context, userID string) (string, error) {
	return us.repo.GetCity(ctx, userID)
}
