package handlers

import (
	"github.com/scrumptious/weather-service/data"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Player struct {
	l *logrus.Logger
}

func NewPlayer() *Player {
	return &Player{}
}

func (p *Player) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	pl := data.GetPlayers()
	//for _, v := range pl {
	//	fmt.Printf("%#v", *v)
	//}
	rw.Write([]byte(pl.ToJSON()))

}
