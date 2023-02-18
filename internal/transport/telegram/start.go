package telegram

import "gopkg.in/telebot.v3"

func (h *Handler) start(c telebot.Context) error {
	shareLoc := telebot.ReplyButton{

		Text: "Поделиться локацией?",

		Location: true,
	}

	repl := &telebot.ReplyMarkup{

		ReplyKeyboard: [][]telebot.ReplyButton{

			[]telebot.ReplyButton{shareLoc},
		},

		OneTimeKeyboard: true,
	}
	return c.Send("Привет, поделитесь своей локацией чтобы мы могли сохранить ваш город в базе данныхы.", repl)
}
