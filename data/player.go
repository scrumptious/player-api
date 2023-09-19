package data

import (
	"encoding/json"
	"log"
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
	Name           string    `json:"name"`
	Level          int       `json:"level"`
	AccountCreated time.Time `json:"accountCreated"`
	Profession     Vocation  `json:"profession"`
}

func (p Player) ToJSON() string {
	//fmt.Printf("type of Player = %T", *p)
	j, err := json.Marshal(p)
	if err != nil {
		log.Fatalln("failed encoding json")
	}
	return string(j)
}

type Players []*Player

func (p *Players) ToJSON() string {
	j, err := json.Marshal(*p)
	if err != nil {
		log.Fatalln("failed encoding json")
	}
	return string(j)
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
		Name:           "Eldernicus",
		Level:          315,
		AccountCreated: time.Date(2015, 8, 13, 12, 23, 5, 0, time.UTC),
		Profession:     ElderDruid,
	},
	&Player{
		Name:           "Magicka",
		Level:          54,
		AccountCreated: time.Date(2013, 8, 13, 12, 23, 5, 0, time.UTC),
		Profession:     MasterSorcerer,
	},
	&Player{
		Name:           "TankEvans",
		Level:          182,
		AccountCreated: time.Date(2014, 8, 13, 12, 23, 5, 0, time.UTC),
		Profession:     EliteKnight,
	},
}
