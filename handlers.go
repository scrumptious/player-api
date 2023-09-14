package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"time"
)

type WeatherHandler struct {
	Url      string
	CacheKey string
}

// WeatherHandler calls weather API to get forecast for tomorrow.
func (wh WeatherHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	start := time.Now()
	locationID := r.URL.Query().Get("locationID")
	wh.CacheKey = fmt.Sprintf("%s-%s", wh.CacheKey, locationID)

	// try to read cached value from Redis, if successful return
	val, err := Rdb.Get(ctx, wh.CacheKey).Result()
	switch {
	case err == redis.Nil:
		log.Println("Cache: key does not exist")
		break
	case err != nil:
		log.Println("Cache: Get failed", err)
		break
	default:
		log.Printf("WeatherService read from cache & took %v\n", time.Since(start))
		_, _ = rw.Write([]byte(val))
		return
	}

	// value not found in cache, make API call to get data
	respCH := make(chan Response)
	ws := NewWeatherService(fmt.Sprintf("%s&res=daily", wh.Url))
	ctxWT, cancel := context.WithTimeout(ctx, time.Millisecond*500)
	ctxVal := context.WithValue(ctxWT, "locationID", locationID)

	defer cancel()

	go func() {
		for {
			w, err := ws.GetWeather(ctxVal)
			if err != nil {
				respCH <- Response{
					Code: 400,
					Msg:  "WeatherHandler - Failed to obtain weather info",
				}
			}
			if w == "" {
				respCH <- Response{
					Code: 404,
					Msg:  "WeatherHandler - weather info not found",
				}
			}
			respCH <- Response{
				Code: 200,
				Msg:  w,
			}
		}
	}()

	select {
	case <-ctxWT.Done():
		rw.Write([]byte(fmt.Sprintf("%v: WeatherService took to long to obtain weather info (%v)", time.Now(), time.Since(start))))
		break
	case res := <-respCH:
		if res.Code != 200 {
			_, _ = rw.Write([]byte(res.Msg))
			return
		}
		Rdb.Set(ctx, wh.CacheKey, res.Msg, App.Config.RedisExp)
		log.Printf("Stored value in cache under '%s' key\n", wh.CacheKey)
		log.Printf("WeatherService took %v to 'GetWeather'", time.Since(start))
		_, _ = rw.Write([]byte(res.Msg))
	}
}
