package JsonData

type WeatherData struct {
	Address           string     `json:"address"`
	Description       string     `json:"description"`
	CurrentConditions CurrentCon `json:"currentConditions"`
	Days              []Days     `json:"days"`
}
type CurrentCon struct {
	Time      string  `json:"datetime"`
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feelslike"`
	Humidity  float64 `json:"humidity"`
	Sunrise   string  `json:"sunrise"`
	Sunset    string  `json:"sunset"`
	Pressure  float64 `json:"pressure"`
}
type Days struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feelslike"`
	Humidity  float64 `json:"humidity"`
	Sunrise   string  `json:"sunrise"`
	Sunset    string  `json:"sunset"`
	Pressure  float64 `json:"pressure"`
	TempMax   float64 `json:"tempmax"`
	TempMin   float64 `json:"tempmin"`
	Hours     []Hours `json:"hours"`
}
type Hours struct {
	Temp       float64 `json:"temp"`
	Conditions string  `json:"conditions"`
}

type Data struct {
	Address     string
	Description string
	Time        string
	Temp        string
	FeelsLike   string
	Humidity    string
	Sunrise     string
	Sunset      string
	Pressure    string
	TempMax     string
	TempMin     string
}
