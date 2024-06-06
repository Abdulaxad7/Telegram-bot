package router

import (
	"TelegramBot/cmd/controller"
	"gopkg.in/telebot.v3"
	"log"
)

var bot *telebot.Bot
var Check = true

func Start() {

	var err error
	bot, err = controller.Config()
	if err != nil {
		log.Fatal(err)
	}
	bot.Handle("/start", LangHandler)

	log.Println("Bot started")
	bot.Start()

}

func LangHandler(c telebot.Context) error {
	var in InlineButton
	var err error
	username := c.Message().Sender.FirstName

	if Check {
		photo := &telebot.Photo{
			File: telebot.FromURL("https://www.hdwallpapers.in/download/green_forest_mountain_with_mist_during_morning_time_4k_hd_nature-3840x2160.jpg"),
		}
		err = c.Send(photo)

	}

	err = c.Delete()
	if err != nil {
		log.Fatal(err)
	}
	return c.Send("Wahoooo. Hi "+username+""+"\nPlease select a language.", &telebot.SendOptions{
		ReplyMarkup: &telebot.ReplyMarkup{
			InlineKeyboard: in.LanguageSelection(bot, c),
		},
	})
}

func CountrySel(c telebot.Context) error {
	var in InlineButton

	var err error
	err = c.Delete()
	if err != nil {
		log.Fatal(err)
	}
	return c.Send("Let's select your Country.", &telebot.SendOptions{
		ReplyMarkup: &telebot.ReplyMarkup{
			InlineKeyboard: in.CountrySelection(bot),
		},
	})
}

func BaseClimate(c telebot.Context) error {
	var k KeyboardButton

	var err error
	err = c.Delete()
	if err != nil {
		log.Fatal(err)
	}
	return c.Send("Choose a weather option", &telebot.SendOptions{
		ReplyMarkup: &telebot.ReplyMarkup{
			ResizeKeyboard: true,
			ReplyKeyboard:  k.WeatherSel(bot),
		},
	})
}
