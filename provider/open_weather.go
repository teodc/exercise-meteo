package provider

import (
	"encoding/json"
	"net/http"
)

// OpenWeatherKey constant
const OpenWeatherKey string = "open_weather"

// OpenWeatherProvider struct
type OpenWeatherProvider struct {
	APIKey string
	Format string
	Units  string
}

// Build provider method
func (provider OpenWeatherProvider) Build(apiKey string, format string, units string) OpenWeatherProvider {
	return OpenWeatherProvider{
		APIKey: apiKey,
		Format: format,
		Units:  units,
	}
}

// Temperature provider method
func (provider OpenWeatherProvider) Temperature(city string) (float64, error) {
	uri := "http://api.openweathermap.org/data/2.5/weather?APPID=" + provider.APIKey + "&q=" + city + "&mode=" + provider.Format + "&units=" + provider.Units

	response, error := http.Get(uri)

	if error != nil {
		return 0.0, error
	}

	defer response.Body.Close()

	var data struct {
		Main struct {
			Temperature float64 `json:"temp"`
		} `json:"main"`
	}

	if error := json.NewDecoder(response.Body).Decode(&data); error != nil {
		return 0.0, error
	}

	//log.Println(openWeatherKey, city, data.Main.Temperature)

	return data.Main.Temperature, nil
}
