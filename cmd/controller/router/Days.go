package router

import (
	"TelegramBot/cmd/JsonData"
	"TelegramBot/cmd/WeatherApi"
	"gopkg.in/telebot.v3"
	"log"
)

func (k KeyboardButton) today(c telebot.Context) error {
	get, err := getDataAsString(0)
	if err != nil {
		err = c.Send("Ooops🌪️, We are sorry, something went wrong on our server")
		log.Fatal(err)
	}
	return c.Send(get+"\n", &telebot.SendOptions{
		ReplyMarkup: &telebot.ReplyMarkup{
			ResizeKeyboard: true,
			ReplyKeyboard:  k.WeatherSel(bot),
		},
	})
}
func (k KeyboardButton) tomorrow(c telebot.Context) error {
	get, err := getDataAsString(1)
	if err != nil {
		err = c.Send("Ooops🌪️, We are sorry, something went wrong on our server")
		log.Fatal(err)
	}
	return c.Send(get+"\n", &telebot.SendOptions{
		ReplyMarkup: &telebot.ReplyMarkup{
			ResizeKeyboard: true,
			ReplyKeyboard:  k.WeatherSel(bot),
		},
	})
}
func (k KeyboardButton) afterTomorrow(c telebot.Context) error {
	get, err := getDataAsString(2)

	if err != nil {
		err = c.Send("Ooops🌪️, We are sorry, something went wrong on our server")
		log.Fatal(err)
	}

	return c.Send(get+"\n", &telebot.SendOptions{
		ReplyMarkup: &telebot.ReplyMarkup{
			ResizeKeyboard: true,
			ReplyKeyboard:  k.WeatherSel(bot),
		},
	})
}

func getDataAsString(check int) (string, error) {
	var err error
	var data *JsonData.Data
	info, err := WeatherApi.Info()
	if check == 0 {
		data, err = WeatherApi.Today()

	} else if check == 1 {
		data, err = WeatherApi.Tomorrow()
	} else {
		data, err = WeatherApi.AfterTomorrow()
	}
	if err != nil {
		return "", err
	}
	return "\nTemperature of:    " + info.Address +
		"\n\n🌡️Temperature is gonna be:    " + data.Temp + "º" +
		"\n🔍Feels Like:    " + data.FeelsLike + "º" +
		"\n\n⏱️Humidity:    " + data.Humidity + "%" +
		"\n🌊Pressure:    " + data.Pressure + "hPa" +
		"\n\n⬆️Maximum Temperature:    " + data.TempMax + "º" +
		"\n⬇️Minimum Temperature:    " + data.TempMin + "º" +
		"\n\n🌞Sunrise:    " + data.Sunrise + "AM" +
		"\n🌝Sunset:    " + data.Sunset + "PM", nil
}
