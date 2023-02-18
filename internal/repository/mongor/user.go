package mongor

import (
	"Levap123/weather-bot/internal/domain"
	"Levap123/weather-bot/internal/entity"
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	coll *mongo.Collection
}

func NewUser(client *mongo.Client) *UserRepo {
	return &UserRepo{
		coll: client.Database("tg").Collection("users"),
	}
}

func (ur *UserRepo) Create(ctx context.Context, user *entity.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := ur.coll.InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("mogo repo - create user - %w", err)
	}
	return nil
}

func (ur *UserRepo) GetCity(ctx context.Context, userID int64) (string, error) {
	filter := bson.M{
		"id": userID,
	}
	result := ur.coll.FindOne(ctx, filter)
	var user entity.User
	if err := result.Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return "", fmt.Errorf("mongo repo - get user - %w", domain.ErrUserNotFound)
		}
		return "", fmt.Errorf("mongo repo - get user - %w", err)
	}
	return user.City, nil
}
