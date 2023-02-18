package telegram

import (
	"Levap123/weather-bot/internal/entity"
	"fmt"

	"gopkg.in/telebot.v3"
)

func (h *Handler) saveUser(c telebot.Context) {
	userID := c.Sender().ID
	fmt.Println(userID)
	user := &entity.User{
		ID: userID,
		
	}
	h.service.UserService.Create()
}
