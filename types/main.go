package types

import (
	"github.com/scrumptious/weather-service/config"
)

var App *config.Application

type Response struct {
	Code int
	Msg  string
}
