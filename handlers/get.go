package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/scrumptious/weather-service/data"
	"net/http"
	"strconv"
)

// swagger:route GET /player/{id} players getPlayer
//
// Returns a requested Player from a data store
//
// Produces:
// - application/json
//
// Parameters:
// + name: id
//   in: query
//   description: ID of player to get
//   required: true
//	 type: integer
//
// Responses:
// 200: playerResponse

// GetPlayer returns a requested Player from a data store
func (p *Player) GetPlayer(rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(rw, fmt.Sprintf(`{"message": "incorrect or missing player id", "error": "%s"}`), http.StatusBadRequest)
	}
	player := data.GetPlayer(id)
	if player == nil {
		http.Error(rw, "player with given ID not found", http.StatusBadRequest)
	}
	player.WriteToJSON(rw)
}

// swagger:route GET /players players listPlayers
//
// Returns a list of Players from data store
//
// Produces:
// - application/json
//
// Responses:
// 200: playersResponse

// GetPlayers returns a list of all players from data source
func (p *Player) GetPlayers(rw http.ResponseWriter, r *http.Request) {
	pl := data.GetPlayers()
	pl.WriteToJSON(rw)
}
