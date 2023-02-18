package telegram

import (
	"gopkg.in/telebot.v3"
)

func (h *Handler) InitRoutes() {
	h.b.Handle(telebot.OnLocation, h.saveUser)
	h.b.Handle("/weather", h.getWeatherByCity)
	h.b.Handle("/save", h.saveUser)
	h.b.Handle("/start", h.start)
}
