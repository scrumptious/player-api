package handlers

import (
	"fmt"
	"github.com/scrumptious/weather-service/internal/data"
	"net/http"
)

// swagger:route POST /player players addPlayer
//
// Adds a Player to the data store
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// Responses:
// 200: playersResponse

// PostPlayer adds a Player to the data store
func (p *Player) PostPlayer(rw http.ResponseWriter, r *http.Request) {
	id := getNextID()
	pl := r.Context().Value(PlayerKey{}).(*data.Player)
	pl.ID = id

	err := pl.Validate()
	if err != nil {
		http.Error(rw, fmt.Sprintf(`{"message": "player failed validation", "error": "%s"}`, err), http.StatusBadRequest)
		return
	}
	data.AddPlayer(pl)
	pls := data.GetPlayers()
	pls.WriteToJSON(rw)
}
