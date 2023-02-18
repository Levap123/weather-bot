package telegram

import (
	"context"
	"fmt"
	"time"

	"gopkg.in/telebot.v3"
)

func (h *Handler) saveUser(c telebot.Context) error {
	lat := c.Message().Location.Lat
	lng := c.Message().Location.Lng
	userID := c.Sender().ID
	fmt.Println(userID)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()
	if err := h.service.UserService.Create(ctx, userID, lng, lat); err != nil {
		h.lg.Errorln(err)
		return err
	}
	return c.Send(`вы добавлены в нашу базу данных! Теперь можете смотреть погоду в своем городе без указания города!
Если вы переехали или находитесь в другом городе, то просто пропишите команду еще раз`)
}
