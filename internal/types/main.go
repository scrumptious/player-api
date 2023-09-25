package types

import (
	"github.com/scrumptious/weather-service/internal/config"
)

var App *config.Application

type Response struct {
	Code int
	Msg  string
}
