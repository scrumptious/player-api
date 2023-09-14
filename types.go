package main

import "github.com/redis/go-redis/v9"

var Rdb *redis.Client
var App *Application

type Response struct {
	Code int
	Msg  string
}
type WeatherData struct {
	Day           string `json:"day"`
	Temperature   int    `json:"temperature"`
	Imperial      bool   `json:"imperial"`
	WindSpeed     int    `json:"wind_speed"`
	WindDirection string `json:"wind_direction"`
	Humidity      int    `json:"humidity"`
	MaxUV         string `json:"max_uv"`
}

type Param struct {
	Name  string `json:"name"`
	Units string `json:"units"`
}

type Period struct {
	Type  string `json:"type"`
	Value string `json:"value"`
	Rep   []Rep
}

type Rep struct {
	D  string
	Dm string
	Hn string
	U  string
	S  string
}

type WeatherAPIData struct {
	SiteRep struct {
		Wx struct {
			params []Param
		}
		DV struct {
			DataDate string `json:"data_date"`
			Type     string `json:"type"`
			Location struct {
				name      string
				country   string
				elevation float32
				Period    []Period
			}
		}
	}
}
