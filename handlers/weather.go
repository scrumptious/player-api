package handlers

import (
	"context"
	"fmt"
	red "github.com/redis/go-redis/v9"
	"github.com/scrumptious/weather-service/data"
	"github.com/scrumptious/weather-service/redis"
	"github.com/scrumptious/weather-service/service"
	"github.com/scrumptious/weather-service/types"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// Weather is a representation of weather handler
type Weather struct {
	l        *logrus.Logger
	url      string
	cacheKey string
	cached   bool
}

// NewWeather is constructor for weather handler
func NewWeather(url, cacheKey string, l *logrus.Logger) *Weather {
	return &Weather{url: url, cacheKey: cacheKey, l: l}
}

// GetWeather will call weather API to get forecast for tomorrow and return JSON
func (w *Weather) GetWeather(rw http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	start := time.Now()
	locationID := r.URL.Query().Get("locationID")
	w.cacheKey = fmt.Sprintf("%s-%s", w.cacheKey, locationID)

	// try to read cached value from Redis, if successful return
	val, err := redis.R.Get(ctx, w.cacheKey).Result()
	switch {
	case err == red.Nil:
		//do nothing here, get a fresh weather data from API
		break
	case err != nil:
		w.l.WithFields(logrus.Fields{
			"error": err,
		}).Fatalln("failed to read from cache")
		break
	default:
		w.cached = true
		w.l.WithFields(logrus.Fields{
			"response_time": fmt.Sprintf("%vms", time.Since(start).Milliseconds()),
			"cached":        w.cached,
		}).Println("reading from cache took")
		_, _ = rw.Write([]byte(val))
		return
	}

	// value not found in cache, make API call to get weather data
	w.cached = false
	respCH := make(chan types.Response)
	ws := service.NewWeatherService(fmt.Sprintf("%s&res=daily", w.url))
	ctxWT, cancel := context.WithTimeout(ctx, time.Millisecond*500)
	ctxVal := context.WithValue(ctxWT, "locationID", locationID)

	defer cancel()

	go func() {
		for {
			wd, err := ws.GetWeather(ctxVal)
			if err != nil {
				respCH <- types.Response{
					Code: http.StatusBadRequest,
					Msg:  fmt.Sprintf("failed to obtain weather info\n%s", err),
				}
			}

			res, err := wd.ToJson()
			if err != nil {
				respCH <- types.Response{
					Code: http.StatusInternalServerError,
					Msg:  err.Error(),
				}
			} else {
				respCH <- types.Response{
					Code: http.StatusOK,
					Msg:  res,
				}
			}
		}
	}()

	select {
	case <-ctxWT.Done():
		w.l.WithFields(logrus.Fields{
			"response_time": fmt.Sprintf("%vms", time.Since(start).Milliseconds()),
			"cached":        w.cached,
		}).Fatalln("it took to long to obtain weather info")
		_, _ = rw.Write([]byte("it took to long to obtain weather info"))
		break
	case res := <-respCH:
		if res.Code != 200 {
			_, _ = rw.Write([]byte(res.Msg))
			return
		}

		redis.R.Set(ctx, w.cacheKey, res.Msg, types.App.Config.RedisExp)
		w.l.WithFields(logrus.Fields{
			"response_time": fmt.Sprintf("%vms", time.Since(start).Milliseconds()),
			"cached":        w.cached,
		}).Println("Time to get weather")
		err = data.StoreWeather(res.Msg)
		if err != nil {
			w.l.WithFields(logrus.Fields{
				"locationID": locationID,
				"cached":     w.cached,
			}).Error("failed to store weather data for this location")
		}
		_, _ = rw.Write([]byte(res.Msg))
	}
}
