// Package handlers
// Package classification of Player API
//
// # Documentation of Player API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/scrumptious/weather-service/data"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

// Single player returned by the response
// swagger:response playerResponse
type playerResponseDef struct {
	// Single player
	// in: body
	Body Player
}

// A list of Players returned by the response
// swagger:response playersResponse
type playersResponseDef struct {
	// All players
	// in: body
	Body []data.Player
}

// swagger:parameters getPlayer deletePlayer
type idParam struct {
	// in: query
	// minimum: 0
	// pattern: /[0-9]+/
	// unique: true
	// required: true
	id int64
}

type noContentDef struct {
}

type Player struct {
	l *logrus.Logger
}

func NewPlayer() *Player {
	return &Player{}
}

func getNextID() int {
	pl := data.GetPlayers()
	return pl[len(pl)-1].ID + 1
}

type PlayerKey struct{}

func (p *Player) MiddlewareSetIDCheckUniqueName(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		player := &data.Player{}
		err := player.FromJSON(r)
		if err != nil {
			http.Error(rw, fmt.Sprintf(`{"message": "failed to read request body", "error": "%s"}`, err), http.StatusBadRequest)
			return
		}
		vars := mux.Vars(r)
		var id int

		switch r.Method {
		case http.MethodPost:
			id = getNextID()
			break
		case http.MethodPut:
			id, _ = strconv.Atoi(vars["id"])
		}
		player.ID = id
		//check for unique name
		unique := true
		for _, v := range data.GetPlayers() {
			//editing user vs updating user with name that is already taken
			if v.Name == player.Name && v.ID != player.ID {
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

func (p *Player) MiddlewarePopulateLastModified(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		pl := r.Context().Value(PlayerKey{}).(*data.Player)
		pl.UpdatedAt = time.Now().Unix()

		ctx := context.WithValue(r.Context(), PlayerKey{}, pl)
		req := r.WithContext(ctx)
		next.ServeHTTP(rw, req)
	})
}

func MiddlewareValidatePlayer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		player := r.Context().Value(PlayerKey{}).(data.Player)
		err := player.Validate()
		if err != nil {
			http.Error(rw, fmt.Sprintf(`"error": "%s"`, err), http.StatusBadRequest)
		}
		next.ServeHTTP(rw, r)
	})
}

func MiddlewareWithValueTemplate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		p := r.Context().Value(PlayerKey{}).(data.Player)
		p.ID = 2 * p.ID
		ctx := context.WithValue(r.Context(), PlayerKey{}, p)
		req := r.WithContext(ctx)
		next.ServeHTTP(rw, req)
	})
}
