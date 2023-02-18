package entity

import "time"

type User struct {
	ID        int64     `bson:"id"`
	City      string    `bson:"city" json:"city"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}
