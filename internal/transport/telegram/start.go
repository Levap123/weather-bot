package telegram

import "gopkg.in/telebot.v3"

var (
	shareLoc = telebot.ReplyButton{

		Text: "Поделиться локацией?",

		Location: true,
	}

	repl = &telebot.ReplyMarkup{

		ReplyKeyboard: [][]telebot.ReplyButton{

			[]telebot.ReplyButton{shareLoc},
		},

		OneTimeKeyboard: true,
	}
)

func (h *Handler) start(c telebot.Context) error {

	return c.Send("Привет, поделитесь своей локацией чтобы мы могли сохранить ваш город в базе данных.", repl)
}
