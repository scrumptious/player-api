package data

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Vocation int

const (
	None Vocation = iota
	Knight
	EliteKnight
	Sorcerer
	MasterSorcerer
	Paladin
	RoyalPaladin
	Druid
	ElderDruid
)

func (v Vocation) String() string {
	switch v {
	case Knight:
		return "Knight"
	case EliteKnight:
		return "Elite Knight"
	case Sorcerer:
		return "Sorcerer"
	case MasterSorcerer:
		return "Master Sorcerer"
	case Paladin:
		return "Paladin"
	case RoyalPaladin:
		return "Royal Paladin"
	case Druid:
		return "Druid"
	case ElderDruid:
		return "Elder Druid"
	default:
		return "None"
	}
}

type Player struct {
	ID             int       `json:"-"`
	Name           string    `json:"name"`
	Level          int       `json:"level"`
	AccountCreated time.Time `json:"accountCreated"`
	Vocation       Vocation  `json:"vocation"`
}

func (p *Player) ToJSON() string {
	//fmt.Printf("type of Player = %T", *p)
	j, err := json.Marshal(p)
	if err != nil {
		log.Fatalln("failed encoding json")
	}
	return string(j)
}

func FromJSON(r *http.Request) (*Player, error) {
	pl := &Player{}
	err := json.NewDecoder(r.Body).Decode(pl)
	if err != nil {
		return nil, fmt.Errorf("failed to decode json")
	}
	return pl, nil
}

type Players []*Player

func (p *Players) ToJSON() string {
	j, err := json.Marshal(*p)
	if err != nil {
		log.Fatalln("failed encoding json")
	}
	return string(j)
}

func (p *Players) WriteToJson(rw http.ResponseWriter) {
	err := json.NewEncoder(rw).Encode(p)
	if err != nil {
		http.Error(rw, "failed to encode json", http.StatusInternalServerError)
	}
}

//func GetPlayer(name string) *Player {
//	for v, _ := range playersList {
//		if
//	}
//}

func GetPlayers() Players {
	return playersList
}

var playersList = Players{
	&Player{
		ID:             1,
		Name:           "Eldernicus",
		Level:          315,
		AccountCreated: time.Date(2015, 8, 13, 12, 23, 5, 0, time.UTC),
		Vocation:       ElderDruid,
	},
	&Player{
		ID:             2,
		Name:           "Magicka",
		Level:          54,
		AccountCreated: time.Date(2013, 8, 13, 12, 23, 5, 0, time.UTC),
		Vocation:       MasterSorcerer,
	},
	&Player{
		ID:             3,
		Name:           "TankEvans",
		Level:          182,
		AccountCreated: time.Date(2014, 8, 13, 12, 23, 5, 0, time.UTC),
		Vocation:       EliteKnight,
	},
}
