package main

import (
	"fmt"
	"github.com/gorilla/mux"
	red "github.com/redis/go-redis/v9"
	"github.com/scrumptious/weather-service/config"
	"github.com/scrumptious/weather-service/handlers"
	"github.com/scrumptious/weather-service/redis"
	"github.com/scrumptious/weather-service/types"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func initApp() {
	types.App = &config.Application{}
}

var log = logrus.New()

func main() {
	log.SetReportCaller(true)
	log.SetFormatter(&logrus.JSONFormatter{})
	initApp()
	types.App.InitConfig()

	redis.InitRedis(&red.Options{
		Addr:     types.App.Config.RedisAddr,
		Password: "",
		DB:       0})

	wh := handlers.NewWeather(
		types.App.Config.WeatherApiUri+types.App.Config.WeatherApiKey,
		"weatherservice",
		log,
	)
	ph := handlers.NewPing(log)
	plh := handlers.NewPlayer()
	r := mux.NewRouter()

	r.Handle("/ping", ph)
	r.Handle("/weather", wh).Methods("GET")
	r.Handle("/player", plh)

	s := &http.Server{
		Addr:         fmt.Sprintf("localhost:%s", types.App.Config.Port),
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      r,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatalln("Failed to initialize a web server")
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	sig := <-sigChan
	log.Println("Received interrupt, shutting down gracefully", sig)

}
