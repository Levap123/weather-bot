package telegram

import (
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
)

func (h *Handler) InitRoutes() {
	h.b.Use(middleware.Logger())
	h.b.Use(middleware.AutoRespond())
	
	h.b.Handle(telebot.OnLocation, h.saveUser)
	h.b.Handle("/weather", h.getWeatherByCity)
	h.b.Handle(&shareLoc, h.saveUser)
	h.b.Handle("/start", h.start)
	h.b.Handle("/qr", h.getQRByLink)
}
