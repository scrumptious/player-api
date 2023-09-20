package handlers

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

type Ping struct {
	l *logrus.Logger
}

func NewPing(l *logrus.Logger) *Ping {
	return &Ping{l: l}
}

func (p *Ping) Ping(rw http.ResponseWriter, r *http.Request) {
	_, err := rw.Write([]byte("pong"))
	if err != nil {
		p.l.Println("Failed to write response", err)
	}
}
