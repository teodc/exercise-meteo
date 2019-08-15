package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/teodc/meteo/config"
	"github.com/teodc/meteo/provider"
)

type weatherProvider interface {
	Temperature(city string) (float64, error)
}

type weatherProviders map[string]weatherProvider

var (
	conf      = config.Load()
	providers = make(weatherProviders)
)

func init() {
	providers[provider.OpenWeatherKey] = provider.OpenWeatherProvider{
		APIKey: conf.OpenWeather.APIKey,
		Format: conf.OpenWeather.Format,
		Units:  conf.OpenWeather.Units,
	}
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/weather/:city", getWeatherForCity)

	e.Logger.Fatal(e.Start("localhost:8080"))
}

func getWeatherForCity(context echo.Context) error {
	city := context.Param("city")

	temperature, error := providers.averageTemperature(city)

	if error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, error.Error())
	}

	context.Response().Header().Set("Content-Type", "application/json; charset=utf-8")

	return context.JSON(http.StatusOK, map[string]interface{}{
		"city":        city,
		"temperature": temperature,
	})
}

func (providers weatherProviders) averageTemperature(city string) (float64, error) {
	temperatures := make(chan float64, len(providers))
	errors := make(chan error, len(providers))

	for _, provider := range providers {
		go func(provider weatherProvider) {
			temperature, error := provider.Temperature(city)
			if error != nil {
				errors <- error
				return
			}
			temperatures <- temperature
		}(provider)
	}

	var sumOfTemperatures float64

	for i := 0; i < len(providers); i++ {
		select {
		case temperature := <-temperatures:
			sumOfTemperatures += temperature
		case error := <-errors:
			return 0.0, error
		}
	}

	return sumOfTemperatures / float64(len(providers)), nil
}
