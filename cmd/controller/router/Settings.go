package router

import (
	"gopkg.in/telebot.v3"
)

func (k KeyboardButton) settings(c telebot.Context) error {

	return c.Send("Choose option ", &telebot.SendOptions{
		ReplyMarkup: &telebot.ReplyMarkup{
			ResizeKeyboard: true,
			ReplyKeyboard:  k.SettingButtons(bot),
		},
	})
}

func (k KeyboardButton) SettingButtons(bot *telebot.Bot) [][]telebot.ReplyButton {

	buttonCountry := telebot.ReplyButton{Text: "Select Another Country 🏴"}
	buttonThankToDeveloper := telebot.ReplyButton{Text: "Thank the developer 💻"}
	selectLanguage := telebot.ReplyButton{Text: "Select Another Language"}
	buttons := [][]telebot.ReplyButton{
		{buttonCountry, selectLanguage}, {buttonThankToDeveloper},
	}
	Check = false
	bot.Handle(&buttonCountry, CountrySel)
	bot.Handle(&selectLanguage, LangHandler)
	bot.Handle(&buttonThankToDeveloper, k.thankHandler)

	return buttons

}

func (k KeyboardButton) thankHandler(c telebot.Context) error {

	anim := &telebot.Photo{File: telebot.FromURL("https://inhalist.web.app/img/thanks.jpg")}
	c.Send(anim)
	return c.Send("You are Welcome 👍🏻", &telebot.SendOptions{
		ReplyMarkup: &telebot.ReplyMarkup{
			ResizeKeyboard: true,
			ReplyKeyboard:  k.WeatherSel(bot),
		},
	})
}
