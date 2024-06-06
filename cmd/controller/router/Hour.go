package router

import (
	"TelegramBot/cmd/WeatherApi"
	"gopkg.in/telebot.v3"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func (k KeyboardButton) hourly(c telebot.Context) error {
	var err error
	err = c.Send("On what time?")
	var sticker string
	bot.Handle(telebot.OnText, func(c telebot.Context) error {
		time := c.Message().Text
		t, _ := strconv.Atoi(time)
		if len(time) > 2 || !unicode.IsDigit(rune(time[0])) || !unicode.IsDigit(rune(time[1])) || (t < 0 || t >= 24) {
			return c.Send("Ooops🌪️, It seems you entered wrong time", &telebot.SendOptions{
				ReplyMarkup: &telebot.ReplyMarkup{
					ResizeKeyboard: true,
					ReplyKeyboard:  k.WeatherSel(bot),
				},
			})

		}
		var tim string
		s, _ := strconv.Atoi(time)
		temp, con, _ := WeatherApi.Hourly(s)
		if s < 12 {
			tim = "AM"
		} else {
			tim = "PM"
		}
		if con == "Partially cloudy" {
			sticker = "⛅️"
		} else if con == "Cloudy" {
			sticker = "☁️"
		} else if strings.Contains(con, "Rain") {
			sticker = "🌧️"
		} else if strings.Contains(con, "Snow") {
			sticker = "❄️"
		} else {
			sticker = "☀️"
		}

		return c.Send("The temperature at "+time+":00"+tim+" is gonna be:   "+temp+
			"º\nThe condition is "+sticker+":   "+con, &telebot.SendOptions{
			ReplyMarkup: &telebot.ReplyMarkup{
				ResizeKeyboard: true,
				ReplyKeyboard:  k.WeatherSel(bot),
			},
		})
	})
	if err != nil {
		err = c.Send("Ooops🌪️, We are sorry, something went wrong on our server")
		log.Fatal(err)
	}
	return err
}
