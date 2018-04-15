package padArchiver

import (
	"testing"

	"github.com/fatih/color"
)

const checkIcon = "\xE2\x9C\x94 "

func TestAddPad(t *testing.T) {
	color.Blue("TestAddPad")
	repo := OpenRepo("Repo01")
	_, err := repo.StorePad("http://board.net/p/pad1", "Group1", "pad1", true)

	if err == nil {
		color.Green(checkIcon + "checked AddPad")
	} else {
		color.Red(err.Error())
		t.Errorf("Error AddPad")
	}
}
