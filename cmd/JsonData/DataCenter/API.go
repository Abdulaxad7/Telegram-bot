package DataCenter

import (
	"TelegramBot/cmd/JsonData"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"net/http"
)

func API(country string) (*JsonData.WeatherData, error) {
	var apiKey = "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/Country?unitGroup=metric&key=RVGU33YYE4JR9EKFTRF8F8SUF&contentType=json"

	apiKey = strings.Replace(apiKey, "Country", country, 1)

	var data JsonData.WeatherData
	resp, err := http.Get(apiKey)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: status code %d\n", resp.StatusCode)
		return nil, err
	}

	if err = json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func Check(userInput string) bool {

	data, err := API(userInput)
	if err != nil || data == nil {
		return false
	}
	return true
}
