package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/scrumptious/weather-service/data"
	"net/http"
	"strconv"
)

// swagger:route DELETE /player/{id} players deletePlayer
//
// Removes a requested Player from data store
//
// Produces:
// - application/json
//
// Parameters:
// + name: id
//   in: query
//   description: ID of player to delete
//   required: true
//	 type: integer
//
// Responses:
// 200: playersResponse

// DeletePlayer removes requested Player from data store
func (p *Player) DeletePlayer(rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(rw, fmt.Sprintf(`{"message": "incorrect or missing player id", "error": "%s"}`, err), http.StatusBadRequest)
		return
	}

	err = data.DeletePlayer(id)
	if err == data.ErrNotFound {
		http.Error(rw, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, fmt.Sprintf("player not found, err: %s", err), http.StatusInternalServerError)
	}
	pls := data.GetPlayers()
	pls.WriteToJSON(rw)
}
