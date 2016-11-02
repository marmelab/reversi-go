package player

import (
  "testing"
  "reversi/game/cell"
)

func TestNewPlayerShouldReturnNewPlayer(t *testing.T) {

  player := New("doe", BrainTypeHuman, cell.TypeBlack)
  expectedPlayer := Player{"doe", BrainTypeHuman, cell.TypeBlack}

  if player != expectedPlayer {
    t.Error("New doesn't return expected player struct")
  }

}
