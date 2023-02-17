package telegram

func (h *Handler) InitRoutes() {
	h.b.Handle("/weather", h.getWeatherByCity)
}
