package handlers

import (
	"fmt"
	"github.com/scrumptious/weather-service/data"
	"net/http"
)

func (p *Player) PutPlayer(rw http.ResponseWriter, r *http.Request) {
	pls := data.GetPlayers()
	updated := r.Context().Value(PlayerKey{}).(*data.Player)
	if data.FindPlayerWithID(updated.ID) == -1 {
		http.Error(rw, `{"error": "player not found"}`, http.StatusBadRequest)
		return
	}

	err := updated.Validate()
	if err != nil {
		http.Error(rw, fmt.Sprintf(`{"message": "player failed validation", "error": "%s"}`, err), http.StatusBadRequest)
		return
	}
	pls[updated.ID-1] = updated
	pls.WriteToJSON(rw)

}
