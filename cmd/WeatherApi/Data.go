package WeatherApi

import (
	"TelegramBot/cmd/JsonData"
	"TelegramBot/cmd/JsonData/DataCenter"
	"fmt"
)

var Country string

func Info() (*JsonData.Data, error) {
	var Data JsonData.Data
	data, err := DataCenter.API(Country)
	if err != nil {
		return nil, err
	}
	Data.Address = data.Address
	Data.Description = data.Description
	return &Data, nil
}

func CurrentTempData() (*JsonData.Data, error) {
	data, err := DataCenter.API(Country)
	if err != nil {
		return nil, err
	}
	current := data.CurrentConditions
	Temp := fmt.Sprintf("%.1f", current.Temp)
	FeelsLike := fmt.Sprintf("%.1f", current.FeelsLike)
	Humidity := fmt.Sprintf("%.1f", current.Humidity)
	Pressure := fmt.Sprintf("%.0f", current.Pressure)

	return &JsonData.Data{
		Time:      current.Time,
		Temp:      Temp,
		FeelsLike: FeelsLike,
		Humidity:  Humidity,
		Pressure:  Pressure,
	}, nil
}

func Today() (*JsonData.Data, error) {
	data, err := DataCenter.API(Country)
	if err != nil {
		return nil, err
	}
	today := data.Days[0]
	Temp := fmt.Sprintf("%.1f", today.Temp)
	FeelsLike := fmt.Sprintf("%.1f", today.FeelsLike)
	Humidity := fmt.Sprintf("%.1f", today.Humidity)
	Pressure := fmt.Sprintf("%.0f", today.Pressure)
	TempMax := fmt.Sprintf("%.1f", today.TempMax)
	TempMin := fmt.Sprintf("%.1f", today.TempMin)
	return &JsonData.Data{
		Temp:      Temp,
		FeelsLike: FeelsLike,
		Humidity:  Humidity,
		Pressure:  Pressure,
		TempMax:   TempMax,
		TempMin:   TempMin,
		Sunrise:   today.Sunrise,
		Sunset:    today.Sunset,
	}, nil
}

func Tomorrow() (*JsonData.Data, error) {
	data, err := DataCenter.API(Country)
	if err != nil {
		return nil, err
	}
	tomorrow := data.Days[1]
	Temp := fmt.Sprintf("%.1f", tomorrow.Temp)
	FeelsLike := fmt.Sprintf("%.1f", tomorrow.FeelsLike)
	Humidity := fmt.Sprintf("%.1f", tomorrow.Humidity)
	Pressure := fmt.Sprintf("%.0f", tomorrow.Pressure)
	TempMax := fmt.Sprintf("%.1f", tomorrow.TempMax)
	TempMin := fmt.Sprintf("%.1f", tomorrow.TempMin)
	return &JsonData.Data{
		Temp:      Temp,
		FeelsLike: FeelsLike,
		Humidity:  Humidity,
		Pressure:  Pressure,
		TempMax:   TempMax,
		TempMin:   TempMin,
		Sunrise:   tomorrow.Sunrise,
		Sunset:    tomorrow.Sunset,
	}, nil
}

func AfterTomorrow() (*JsonData.Data, error) {
	data, err := DataCenter.API(Country)
	if err != nil {
		return nil, err
	}
	afterTomorrow := data.Days[2]
	Temp := fmt.Sprintf("%.1f", afterTomorrow.Temp)
	FeelsLike := fmt.Sprintf("%.1f", afterTomorrow.FeelsLike)
	Humidity := fmt.Sprintf("%.1f", afterTomorrow.Humidity)
	Pressure := fmt.Sprintf("%.0f", afterTomorrow.Pressure)
	TempMax := fmt.Sprintf("%.1f", afterTomorrow.TempMax)
	TempMin := fmt.Sprintf("%.1f", afterTomorrow.TempMin)
	return &JsonData.Data{
		Temp:      Temp,
		FeelsLike: FeelsLike,
		Humidity:  Humidity,
		Pressure:  Pressure,
		TempMax:   TempMax,
		TempMin:   TempMin,
		Sunrise:   afterTomorrow.Sunrise,
		Sunset:    afterTomorrow.Sunset,
	}, nil
}

func Hourly(t int) (string, string, error) {
	data, err := DataCenter.API(Country)

	if err != nil {
		return "", "", err
	}
	for i, v := range data.Days {
		if i == 0 {
			for j, k := range v.Hours {
				if j == t {
					temp := fmt.Sprintf("%.1f", k.Temp)
					return temp, k.Conditions, nil
				}
			}
		}
	}
	return "", "", nil
}
