package main

import "time"

type Config struct {
	Port          string
	RedisAddr     string
	WeatherApiUri string
	WeatherApiKey string
	RedisExp      time.Duration
}

type Application struct {
	Config Config
}

func (a *Application) initConfig() {
	a.Config = Config{
		Port:          "40400",
		RedisAddr:     "localhost:6379",
		WeatherApiUri: "http://datapoint.metoffice.gov.uk/public/data/val/wxfcs/all/json/locationID?key=",
		WeatherApiKey: "3768a301-4afa-4038-8ce0-c1eacf4207a4",
		RedisExp:      time.Hour,
	}
}
