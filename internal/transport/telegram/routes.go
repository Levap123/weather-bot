package telegram

func (h *Handler) InitRoutes() {
	h.b.Handle("/weather", h.getWeatherByCity)
	h.b.Handle("/save", h.saveUser)
}
