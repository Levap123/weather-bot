package main

import (
	"Levap123/weather-bot/configs"
	"Levap123/weather-bot/internal/transport/telegram"

	"github.com/Levap123/utils/lg"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	if err := configs.InitConfigs(); err != nil {
		logrus.Fatalln(err)
	}
	logger, err := lg.NewLogger()
	if err != nil {
		logrus.Fatalln(err)
	}
	service := telegram.NewService()
	h, err := telegram.InitBot(viper.GetString("telegram_api"), service, logger)
	if err != nil {
		logrus.Fatalln(err)
	}

	h.InitRoutes()
	h.RunBot()
}
