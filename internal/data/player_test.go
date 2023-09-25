package data

import "testing"

func TestCheckValidation(t *testing.T) {
	player := &Player{
		Name: "",
	}
	err := player.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
