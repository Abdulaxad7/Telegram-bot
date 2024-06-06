package router

import (
	"TelegramBot/cmd/WeatherApi"
	"gopkg.in/telebot.v3"
	"log"
	time2 "time"
)

func (k KeyboardButton) WeatherSel(bot *telebot.Bot) [][]telebot.ReplyButton {
	currentButton := telebot.ReplyButton{Text: "Current Temperature ☀️"}
	todayButton := telebot.ReplyButton{Text: "Today's Temperature 🌤"}
	tomorrowButton := telebot.ReplyButton{Text: "Tomorrow's Temperature ⛅️"}
	afterTomorrowButton := telebot.ReplyButton{Text: "After Tomorrow's Temperature ⛈️"}
	hourlyButton := telebot.ReplyButton{Text: "Hourly Temperature 🌨️"}
	settingsButton := telebot.ReplyButton{Text: "Settings ⚙️"}
	var weatherData [][]telebot.ReplyButton
	weatherData = append(weatherData, []telebot.ReplyButton{currentButton})
	weatherData = append(weatherData, []telebot.ReplyButton{todayButton, tomorrowButton})
	weatherData = append(weatherData, []telebot.ReplyButton{afterTomorrowButton, hourlyButton})
	weatherData = append(weatherData, []telebot.ReplyButton{settingsButton})

	bot.Handle(&currentButton, k.current)
	bot.Handle(&todayButton, k.today)
	bot.Handle(&tomorrowButton, k.tomorrow)
	bot.Handle(&afterTomorrowButton, k.afterTomorrow)
	bot.Handle(&hourlyButton, k.hourly)
	bot.Handle(&settingsButton, k.settings)

	return weatherData
}

func (k KeyboardButton) current(c telebot.Context) error {
	now := time2.Now()
	date := now.Format("02-01-2006")
	info, err := WeatherApi.Info()
	data, err := WeatherApi.CurrentTempData()
	if err != nil {
		err = c.Send("Ooops🌪️, We are sorry, something went wrong on our server")
		log.Fatal(err)
	}
	var timeC string
	f1 := data.Time[0]
	f2 := data.Time[1]
	if (f1 == 50 || f2 >= 50) && (f1 != 48) {
		timeC = data.Time + "PM"
	} else {
		timeC = data.Time + "AM"
	}

	return c.Send("Today is "+date+
		"\n\nWeather of "+info.Address+
		"\n🗒️Description:    "+info.Description+
		"\n\n🕰️Current Time is:    "+timeC+
		"\n🌡️Current Temperature is:    "+data.Temp+"º"+
		"\n🔍Feels Like:    "+data.FeelsLike+"º"+
		"\n⏱️Humidity:    "+data.Humidity+"%"+
		"\n🌊Pressure is:    "+data.Pressure+"hPa",
		&telebot.SendOptions{
			ReplyMarkup: &telebot.ReplyMarkup{
				ResizeKeyboard: true,
				ReplyKeyboard:  k.WeatherSel(bot),
			},
		})
}
