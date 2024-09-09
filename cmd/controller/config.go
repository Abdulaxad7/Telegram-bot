package controller

import (
	"log"
	"os"
	"time"

	"gopkg.in/telebot.v3"
)

func Config() (*telebot.Bot, error) {
	// Replace with your bot token
	botToken := os.Getenv("BOT_TOKEN")

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

//
//msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
////button:= tgbotapi.KeyboardButton{Text: "this is button"}
//tgbotapi.NewKeyboardButtonLocation("df")
//buttons := tgbotapi.NewKeyboardButtonRow(tgbotapi.NewKeyboardButton("Button 1"), tgbotapi.NewKeyboardButton("Button 2"))
////butto:=tgbotapi.NewInlineKeyboardRow( tgbotapi.InlineKeyboardButton{Text: "button"})
//d := tgbotapi.NewReplyKeyboard(buttons)
//msg.ReplyMarkup = d
//tgbotapi.NewReplyKeyboard(buttons)
