package main

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
)

func initApp() {
	App = &Application{}
}

func main() {
	initApp()
	App.initConfig()

	Rdb = redis.NewClient(&redis.Options{
		Addr:     App.Config.RedisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	http.Handle("/weather", WeatherHandler{
		Url:      App.Config.WeatherApiUri + App.Config.WeatherApiKey,
		CacheKey: "weatherservice"})
	err := http.ListenAndServe(fmt.Sprintf(":%s", App.Config.Port), nil)
	if err != nil {
		log.Fatalln("Failed to initialise a web server")
	} else {
		log.Println("Web server started..")
	}
}
