package telegram

import tele "gopkg.in/telebot.v3"

type Handler struct {
	b *tele.Bot
}

func InitBot(token string) (*Handler, error) {
	pref := tele.Settings{
		Token: token,
	}
	b, err := tele.NewBot(pref)
	if err != nil {
		return nil, err
	}
	return &Handler{
		b: b,
	}, nil
}

func (h *Handler) InitRoutes() {
	h.b.Handle("/hello",h.sayHello)
}