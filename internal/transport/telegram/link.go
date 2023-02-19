package telegram

import (
	"github.com/skip2/go-qrcode"
	"gopkg.in/telebot.v3"
)

func (h *Handler) getQRByLink(c telebot.Context) error {
	link := c.Message().Payload
	if link == "" {
		return c.Send(`Например: /qr https://example.org`)
	}

	err := qrcode.WriteFile(link, qrcode.Medium, 256, "qr.png")
	if err != nil {
		h.lg.Errorln(err)
		c.Send(`Извиняемся, ошибка на сервере`)
		return err
	}
	
	qr := &telebot.Photo{File: telebot.FromDisk("qr.png")}
	return c.Send(qr)
}
