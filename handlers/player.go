package handlers

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/scrumptious/weather-service/data"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type Player struct {
	l *logrus.Logger
}

func NewPlayer() *Player {
	return &Player{}
}

func findPlayerWithID(id int) *data.Player {
	for _, v := range data.GetPlayers() {
		if v.ID == id {
			return v
		}
	}
	return nil
}

func (p *Player) GetPlayers(rw http.ResponseWriter, r *http.Request) {
	pl := data.GetPlayers()
	pl.WriteToJson(rw)
}

func (p *Player) PostPlayer(rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(rw, "failed to convert ID", http.StatusInternalServerError)
		return
	}
	pls := data.GetPlayers()
	pl, err := data.FromJSON(r)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	if findPlayerWithID(id) == nil {
		pls[id] = pl
	} else {
		http.Error(rw, "player with this ID already exist", http.StatusBadRequest)
		return
	}

	pls.WriteToJson(rw)
}

type PlayerKey struct{}

func (p *Player) PutPlayer(rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(rw, `{"error": "failed to convert ID"}`, http.StatusInternalServerError)
		return
	}
	pls := data.GetPlayers()
	if findPlayerWithID(id) == nil {
		http.Error(rw, `{"error": "player not found"}`, http.StatusNotFound)
		return
	}
	new := r.Context().Value(PlayerKey{}).(*data.Player)
	if err != nil {
		http.Error(rw, `{"error": "incorrect player data"}`, http.StatusBadRequest)
		return
	}
	new.ID = id
	pls[id-1] = new
	rw.WriteHeader(http.StatusOK)
	pls.WriteToJson(rw)

}

func (p *Player) MiddlewareUniqueName(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		player, err := data.FromJSON(r)
		if err != nil {
			http.Error(rw, "failed to read request body", http.StatusBadRequest)
			return
		}
		unique := true
		for _, v := range data.GetPlayers() {
			if v.Name == player.Name {
				unique = false
			}
		}
		if !unique {
			http.Error(rw, fmt.Sprintf(`{"error": "player with this name already exist"}`), http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), PlayerKey{}, player)
		req := r.WithContext(ctx)
		next.ServeHTTP(rw, req)
	})
}
