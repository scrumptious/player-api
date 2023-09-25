package data

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"regexp"
	"time"
)

type Vocation int64

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
	InvalidVocation
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
	case None:
		return "None"
	default:
		return "invalid"
	}
}

//swagger:model
type Player struct {
	// the id for this player
	//
	// required: true
	// unique: true
	// min: 1
	ID int `json:"id"`
	// the name of this player
	//
	// required: true
	// unique: true
	Name      string   `json:"name" validate:"required,alpha"`
	Level     int      `json:"level" validate:"required,gte=1,lte=10000"`
	CreatedAt int64    `json:"createdAt" validate:"required"`
	UpdatedAt int64    `json:"updatedAt" validate:"required"`
	DeteledAt int64    `json:"-"`
	Vocation  Vocation `json:"vocation" validate:"vocation"`
	Signature string   `json:"signature" validate:"signature"`
}

func (p *Player) ToJSON() string {
	j, err := json.Marshal(p)
	if err != nil {
		log.Fatalln("failed encoding json")
	}
	return string(j)
}

func (p *Player) FromJSON(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(p)
	if err != nil {
		return fmt.Errorf("failed to decode json")
	}
	return nil
}

func ValidateVocation(fl validator.FieldLevel) bool {
	result := fl.Field().Int()
	if result >= int64(InvalidVocation) || result < 0 {
		return false
	}
	return true
}

func ValidateSignature(fl validator.FieldLevel) bool {
	r := `[a-z]{3}[0-9]{3}-[a-z]{4}-[0-9]{4}`
	re := regexp.MustCompile(r)
	matches := re.FindAllStringSubmatch(fl.Field().String(), -1)

	if len(matches) == 0 {
		return false
	}
	return true
}

func (p *Player) Validate() error {
	v := validator.New(validator.WithRequiredStructEnabled())
	err := v.RegisterValidation("vocation", ValidateVocation)
	err = v.RegisterValidation("signature", ValidateSignature)
	if err != nil {
		fmt.Println("failed registering validation function", err)
	}
	return v.Struct(p)
}

func (p Players) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}

type Players []*Player

func (p Players) ToJSON() string {
	j, err := json.Marshal(p)
	if err != nil {
		log.Fatalln("failed encoding json")
	}
	return string(j)
}

// TODO refactor to remove code duplication below

// WriteToJSON writes JSON representation of Player using response writer
func (p *Player) WriteToJSON(rw http.ResponseWriter) {
	err := json.NewEncoder(rw).Encode(p)
	if err != nil {
		http.Error(rw, "failed to encode json", http.StatusInternalServerError)
	}
}

func (p Players) WriteToJSON(rw http.ResponseWriter) {
	err := json.NewEncoder(rw).Encode(p)
	if err != nil {
		http.Error(rw, "failed to encode json", http.StatusInternalServerError)
	}
}

func GetPlayer(id int) *Player {
	for _, p := range GetPlayers() {
		if p.ID == id {
			return p
		}
	}
	return nil
}
func GetPlayers() Players {
	return playersList
}

func FindPlayerWithID(id int) int {
	for i, v := range GetPlayers() {
		if v.ID == id {
			return i
		}
	}
	return -1
}

func AddPlayer(player *Player) Players {
	playersList = append(playersList, player)
	return playersList
}

var ErrNotFound = errors.New("ErrNotFound: player not found")

func DeletePlayer(id int) error {
	i := FindPlayerWithID(id)
	if i == -1 {
		return ErrNotFound
	}
	pls := GetPlayers()
	plsNew := append(pls[:i], pls[i+1:]...)
	playersList = plsNew

	return nil
}

var playersList = Players{
	&Player{
		ID:        1,
		Name:      "Eldernicus",
		Level:     315,
		CreatedAt: time.Date(2015, 8, 13, 12, 23, 5, 0, time.UTC).Unix(),
		UpdatedAt: time.Now().Unix(),
		Vocation:  ElderDruid,
	},
	&Player{
		ID:        2,
		Name:      "Magicka",
		Level:     54,
		CreatedAt: time.Date(2013, 8, 13, 12, 23, 5, 0, time.UTC).Unix(),
		UpdatedAt: time.Now().Unix(),
		Vocation:  MasterSorcerer,
	},
	&Player{
		ID:        3,
		Name:      "TankEvans",
		Level:     182,
		CreatedAt: time.Date(2014, 8, 13, 12, 23, 5, 0, time.UTC).Unix(),
		UpdatedAt: time.Now().Unix(),
		Vocation:  EliteKnight,
	},
}
