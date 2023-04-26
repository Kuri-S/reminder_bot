package main

import (
	"log"
	"github.com/spf13/viper"
	"./bot"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetDefault("debug", false)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	telegramBot, err := bot.NewTelegramBot()
	if err != nil {
		log.Fatalf("Error creating Telegram bot, %s", err)
	}

	err = telegramBot.Start()
	if err != nil {
		log.Fatalf("Error starting Telegram bot, %s", err)
	}
}

