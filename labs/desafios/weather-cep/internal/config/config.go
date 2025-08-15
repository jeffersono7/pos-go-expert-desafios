package config

import "os"

var (
	ApiKey string
)

func init() {
	key := os.Getenv("WEATHER_API_KEY")
	if key == "" {
		panic("WEATHER_API_KEY is not set")
	}
	ApiKey = key
}
