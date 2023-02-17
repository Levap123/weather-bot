package telegram

func (h *Handler) InitRoutes() {
	h.b.Handle("/hello", h.sayHello)
}
