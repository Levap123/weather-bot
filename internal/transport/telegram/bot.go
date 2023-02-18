package telegram

import (
	"github.com/sirupsen/logrus"
	tele "gopkg.in/telebot.v3"
)

type Handler struct {
	b       *tele.Bot
	service *Service
	lg      *logrus.Logger
}

func NewHandler(token string, service *Service, lg *logrus.Logger) (*Handler, error) {
	pref := tele.Settings{
		Token: token,
	}
	b, err := tele.NewBot(pref)
	if err != nil {
		return nil, err
	}
	return &Handler{
		b:       b,
		service: service,
		lg:      lg,
	}, nil
}

func (h *Handler) RunBot() {
	h.b.Start()
}
