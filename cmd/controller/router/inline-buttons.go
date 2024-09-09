package router

import (
	"TelegramBot/cmd/JsonData/DataCenter"
	"TelegramBot/cmd/WeatherApi"
	"gopkg.in/telebot.v3"
	"log"
)

func (InlineButton) LanguageSelection(bot *telebot.Bot, c Context) [][]telebot.InlineButton {
	_ = c.Delete()
	lang1 := telebot.InlineButton{
		Unique: "btn1",
		Text:   "English 🇺🇸",
	}
	lang2 := telebot.InlineButton{
		Unique: "btn2",
		Text:   "Russian 🇷🇺",
	}
	inlineKeyboard := [][]telebot.InlineButton{
		{lang1, lang2},
	}
	bot.Handle(&lang1, CountrySel)
	bot.Handle(&lang2, CountrySel)
	return inlineKeyboard
}

func (InlineButton) CountrySelection(bot *telebot.Bot) [][]telebot.InlineButton {

	var err error
	buttons := []telebot.InlineButton{
		{Unique: "btn1", Text: "Albania"},
		{Unique: "btn2", Text: "Palestine"},
		{Unique: "btn3", Text: "Bahrain"},
		{Unique: "btn4", Text: "Argentina"},
		{Unique: "btn5", Text: "Angola"},
		{Unique: "btn6", Text: "Armenia"},
		{Unique: "btn7", Text: "Delhi"},
		{Unique: "btn8", Text: "San_Fransisco"},
		{Unique: "btn9", Text: "New_York"},
		{Unique: "btn10", Text: "London"},
		{Unique: "btn11", Text: "Paris"},
		{Unique: "btn12", Text: "Kiev"},
		{Unique: "btn13", Text: "Washington"},
		{Unique: "btn14", Text: "Moscow"},
		{Unique: "btn15", Text: "Tashkent"},
		{Unique: "btn16", Text: "Another"},
	}
	regions := inlineButtons(buttons)
	for i := range buttons {
		btn := buttons[i]
		bot.Handle(&btn, func(c telebot.Context) error {
			if btn.Text == "Another" {
				ifAnother(c)
				err = c.Delete()
			} else {
				WeatherApi.Country = btn.Text
				return BaseClimate(c)
			}
			return nil
		})
	}
	if err != nil {
		log.Fatal(err)
	}
	return regions
}

func inlineButtons(buttons []telebot.InlineButton) [][]telebot.InlineButton {
	regions := [][]telebot.InlineButton{
		{buttons[0]},
	}
	for i := 1; i < len(buttons); i += 2 {
		if i+1 < len(buttons) {
			regions = append(regions, []telebot.InlineButton{buttons[i], buttons[i+1]})
		} else {
			regions = append(regions, []telebot.InlineButton{buttons[i]})
		}
	}
	return regions
}

func ifAnother(c Context) {
	var err error
	err = c.Send("Goood, then let's type it")
	bot.Handle(telebot.OnText, func(c telebot.Context) error {
		userInput := c.Message().Text

		if DataCenter.Check(userInput) {
			WeatherApi.Country = userInput
			return BaseClimate(c)

		} else {
			err = c.Send("Ooops, You typed wrong country 🏳️")
			return CountrySel(c)
		}
	})
	if err != nil {
		log.Fatal(err)
	}
}
