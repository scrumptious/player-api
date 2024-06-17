package data

import (
	"testing"
)

func TestCheckValidation(t *testing.T) {
	player := &Player{
		Name:      "Lud",
		Level:     128,
		CreatedAt: 1718105978,
		UpdatedAt: 1718105978,
		Vocation:  ElderDruid,
		Signature: "xbs386-isye-2074",
	}
	err := player.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
