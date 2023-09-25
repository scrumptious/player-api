package data

import (
	"encoding/json"
	"log"
)

// WeatherData represents weather data format that will be returned from this service
type WeatherData struct {
	LocationID    int    `json:"locationID"`
	Day           string `json:"day"`
	Temperature   int    `json:"temperature"`
	Imperial      bool   `json:"imperial"`
	WindSpeed     int    `json:"windSpeed"`
	WindDirection string `json:"windDirection"`
	Humidity      int    `json:"humidity"`
	MaxUV         string `json:"maxUv"`
}

// ToJson will decode data structure into json string
func (wd *WeatherData) ToJson() (string, error) {
	r, err := json.Marshal(wd)
	return string(r), err
}

// Param represents internal object within WeatherAPIData object
type Param struct {
	Name  string `json:"name"`
	Units string `json:"units"`
}

// Period represents internal object within WeatherAPIData object
type Period struct {
	Type  string `json:"type"`
	Value string `json:"value"`
	Rep   []Rep
}

// Rep represents internal object within WeatherAPIData object
type Rep struct {
	D  string
	Dm string
	Hn string
	U  string
	S  string
}

// WeatherAPIData represents format of the data received back from MetOffice weather API
type WeatherAPIData struct {
	SiteRep struct {
		Wx struct {
			params []Param
		}
		DV struct {
			DataDate string `json:"data_date"`
			Type     string `json:"type"`
			Location struct {
				i      string
				name   string
				Period []Period
			}
		}
	}
	LocationID string
}

// StoreWeather is used to check if weather data has already been stored for a given day
// if not it will attempt to store into persistent storage
func StoreWeather(locationID string) error {

	return nil
}

// GetWeather will get weather data from persistent storage
func GetWeather(locationID string) (*WeatherData, error) {
	return nil, nil
}

// PostWeather will attempt to store weather data into persistent storage
func PostWeather(wd string) (int, error) {
	log.Println(wd)
	return 1, nil
}
