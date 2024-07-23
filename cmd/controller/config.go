package controller

import (
	"log"
	"time"

	"gopkg.in/telebot.v3"
)

func Config() (*telebot.Bot, error) {
	// Replace with your bot token
	botToken := "**************************************************"

	pref := telebot.Settings{
		Token:  botToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}
	bot, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return bot, nil

}
