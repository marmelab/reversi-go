package ai

import (
	// "fmt"
	// "time"
	//"reversi/game/board"
	"reversi/game/cell"
	"reversi/game/game"
	"reversi/game/player"
	"testing"
)

func BenchmarkGetBestCellChange(b *testing.B) {

	playerBlack := player.New("John", true, cell.TypeBlack)
	playerWhite := player.New("Doe", true, cell.TypeWhite)

	party := game.New([]player.Player{playerBlack, playerWhite})

	for n := 0; n < b.N; n++ {
		GetBestCellChange(party, playerBlack, 0, 1)
	}

}
