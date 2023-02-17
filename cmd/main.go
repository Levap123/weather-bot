package main

import (
	"Levap123/weather-bot/internal/transport/telegram"

	"github.com/sirupsen/logrus"
)

func main() {
	h, err := telegram.InitBot("6075879974:AAH1M-GjeEJF2RTOpiki81Au-lFioyMiyEY")
	if err != nil {
		logrus.Fatalln(err)
	}
	h.InitRoutes()
	h.RunBot()
}
