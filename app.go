package main

import "time"

type Config struct {
	WeatherApiUri string
	WeatherApiKey string
	RedisExp      time.Duration
}

type Application struct {
	Config Config
}

func (a *Application) initConfig() {
	a.Config = Config{
		WeatherApiUri: "http://datapoint.metoffice.gov.uk/public/data/val/wxfcs/all/json/locationID?key=",
		WeatherApiKey: "3768a301-4afa-4038-8ce0-c1eacf4207a4",
		RedisExp:      time.Hour,
	}
}
