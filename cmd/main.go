package main

import (
	"Levap123/weather-bot/configs"
	"Levap123/weather-bot/internal/repository/mongor"
	"Levap123/weather-bot/internal/service"
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
	cl, err := mongor.InitDb()
	if err != nil {
		logrus.Fatalln(err)
	}
	storage := service.NewStorageMongo(cl)
	service := telegram.NewService(storage)
	h, err := telegram.NewHandler(viper.GetString("telegram_api"), service, logger)
	if err != nil {
		logrus.Fatalln(err)
	}

	logrus.Println("bot is started")
	h.InitRoutes()
	h.RunBot()
}
