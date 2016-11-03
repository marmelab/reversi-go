package player

import (
	"reversi/game/cell"
	"testing"
)

func TestNewPlayerShouldReturnNewPlayer(t *testing.T) {

	player := New("doe", true, cell.TypeBlack)
	expectedPlayer := Player{"doe", true, cell.TypeBlack}

	if player != expectedPlayer {
		t.Error("New doesn't return expected player struct")
	}

}
