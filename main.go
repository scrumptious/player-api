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

	r := mux.NewRouter()
	wh := handlers.NewWeather(
		types.App.Config.WeatherApiUri+types.App.Config.WeatherApiKey,
		"weatherservice",
		log,
	)
	ph := handlers.NewPing(log)
	plh := handlers.NewPlayer()

	getR := r.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/ping", ph.Ping)
	getR.HandleFunc("/weather", wh.GetWeather)
	getR.HandleFunc("/", plh.GetPlayers)

	postR := r.Methods(http.MethodPost).Subrouter()
	postR.HandleFunc("/", plh.PostPlayer)
	postR.Use(plh.MiddlewareSetIDCheckUniqueName, plh.MiddlewarePopulateLastModified)

	putR := r.Methods(http.MethodPut).Subrouter()
	putR.HandleFunc("/{id:[0-9]+}", plh.PutPlayer)
	putR.Use(plh.MiddlewareSetIDCheckUniqueName, plh.MiddlewarePopulateLastModified)

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
