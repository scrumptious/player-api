package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Service interface {
	GetWeather(ctx context.Context) (string, error)
}

type WeatherService struct {
	url string
}

// NewWeatherService returns new instance of WeatherService.
func NewWeatherService(url string) Service {
	return &WeatherService{url: url}
}

// GetWeather tries to obtain weather forecast for tomorrow by calling
// weather API and returning data as json string.
func (ws *WeatherService) GetWeather(ctx context.Context) (string, error) {
	locationID := ctx.Value("locationID").(string)
	url := strings.Replace(ws.url, "locationID", locationID, 1)

	resp, err := http.Get(fmt.Sprintf("%s&res=daily", url))
	if err != nil {
		log.Fatalf("%v: WeatherService - Failed to obtain weather info\n", time.Now())
	}
	defer resp.Body.Close()

	body := &WeatherAPIData{}
	err = json.NewDecoder(resp.Body).Decode(body)
	if err != nil {
		log.Fatalf("%v: WeatherService - Failed to decode weather api response", time.Now())
	}
	if body.SiteRep.DV.Location.Period == nil {
		return "", fmt.Errorf("no weather info found for this location - %s", locationID)
	}

	var res WeatherData

	dayData := body.SiteRep.DV.Location.Period[0].Rep[0] //Period[0] - first day of the forecast, Rep[0] - day, Rep[1] - night

	res.Temperature, _ = strconv.Atoi(dayData.Dm)
	res.WindSpeed, _ = strconv.Atoi(dayData.S)
	res.Humidity, _ = strconv.Atoi(dayData.Hn)
	res.Day = body.SiteRep.DV.Location.Period[0].Value
	res.WindDirection = dayData.D
	res.MaxUV = dayData.U
	res.Imperial = true
	res.LocationID, _ = strconv.Atoi(locationID)
	//res.LocationID, _ = strconv.Atoi(body.SiteRep.DV.Location.i)

	r, err := json.Marshal(res)
	if err != nil {
		log.Fatalf("%v: WeatherService - Failed to json encode weather api response", time.Now())
	}
	return string(r), nil
}
