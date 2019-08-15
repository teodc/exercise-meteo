package config

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// OpenWeatherConfig struct holding your OpenWeather related configuration
type OpenWeatherConfig struct {
	APIKey string
	Format string
	Units  string
}

// Config struct that holds your all your configuration values
type Config struct {
	Env         string
	Debug       bool
	Locale      string
	OpenWeather OpenWeatherConfig
}

// Load your configuration from your .env file
func Load() *Config {
	if error := godotenv.Load(); error != nil {
		log.Println(".env file not found")
	}

	return &Config{
		Env:    getEnvString("APP_ENV", "local"),
		Debug:  getEnvBool("APP_DEBUG", true),
		Locale: getEnvString("APP_LOCALE", "en"),
		OpenWeather: OpenWeatherConfig{
			APIKey: getEnvString("OPEN_WEATHER_API_KEY", "abcd1234"),
			Format: getEnvString("OPEN_WEATHER_FORMAT", "json"),
			Units:  getEnvString("OPEN_WEATHER_UNITS", "metric"),
		},
	}
}

func getEnvString(key string, defaultValue string) string {
	if valueString, exists := os.LookupEnv(key); exists {
		return valueString
	}

	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	valueString := getEnvString(key, "")

	if valueInt, error := strconv.Atoi(valueString); error == nil {
		return valueInt
	}

	return defaultValue
}

func getEnvBool(key string, defaultValue bool) bool {
	valueString := getEnvString(key, "")

	if valueBool, error := strconv.ParseBool(valueString); error == nil {
		return valueBool
	}

	return defaultValue
}

func getEnvSlice(key string, defaultValue []string) []string {
	separator := ","
	valueString := getEnvString(key, "")

	if valueString == "" {
		return defaultValue
	}

	valueSlice := strings.Split(valueString, separator)

	return valueSlice
}
