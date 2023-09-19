package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/scrumptious/weather-service/data"
	"net/http"
	"strconv"
	"strings"
)

type Service interface {
	GetWeather(ctx context.Context) (*data.WeatherData, error)
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
func (ws *WeatherService) GetWeather(ctx context.Context) (*data.WeatherData, error) {
	locationID := ctx.Value("locationID").(string)
	url := strings.Replace(ws.url, "locationID", locationID, 1)

	resp, err := http.Get(fmt.Sprintf("%s&res=daily", url))
	if err != nil {
		return nil, fmt.Errorf("failed to obtain weather info\n%s", err)
	}
	defer resp.Body.Close()

	body := &data.WeatherAPIData{}
	err = json.NewDecoder(resp.Body).Decode(body)
	if err != nil {
		return nil, fmt.Errorf("failed to decode weather api response")
	}
	if body.SiteRep.DV.Location.Period == nil {
		return nil, fmt.Errorf("no weather info found for this location - %s", locationID)
	}
	var res data.WeatherData

	dayData := body.SiteRep.DV.Location.Period[0].Rep[0] //Period[0] - first day of the forecast, Rep[0] - day, Rep[1] - night

	res.Temperature, _ = strconv.Atoi(dayData.Dm)
	res.WindSpeed, _ = strconv.Atoi(dayData.S)
	res.Humidity, _ = strconv.Atoi(dayData.Hn)
	res.Day = body.SiteRep.DV.Location.Period[0].Value
	res.WindDirection = dayData.D
	res.MaxUV = dayData.U
	res.Imperial = true
	res.LocationID, _ = strconv.Atoi(locationID)

	return &res, nil
}
