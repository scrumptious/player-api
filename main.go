package main

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
)

var port = "40400"

func initApp() {
	App = &Application{}
}

func main() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	initApp()
	App.initConfig()

	http.Handle("/weather", WeatherHandler{
		App.Config.WeatherApiUri + App.Config.WeatherApiKey,
		"weatherservice"})
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatalln("Failed to initialise a web server")
	} else {
		log.Println("Web server started..")
	}
}
