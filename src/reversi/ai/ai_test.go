package ai

import (
	"fmt"
	"reversi/game/board"
	"reversi/game/cell"
	//"reversi/game/game"
	//"reversi/game/player"
	"testing"
)

// func BenchmarkGetBestCellChange(b *testing.B) {
//
// 	playerBlack := player.New("John", true, cell.TypeBlack)
// 	playerWhite := player.New("Doe", true, cell.TypeWhite)
//
// 	party := game.New([]player.Player{playerBlack, playerWhite})
//
// 	for n := 0; n < b.N; n++ {
// 		GetBestCellChangeInTime(party, 0, 1)
// 	}
//
// }

func TestGetBestCellChangeInTime(t *testing.T) {

	currentBoard, _ := board.InitCells(board.New(8, 8))
	fmt.Println(GetBestCellChangeInTime(currentBoard, cell.TypeBlack))

}

// func TestGetZoningScore(t *testing.T) {
//
// 	board := board.Board{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
//
// 	if GetZoningScore([]cell.Cell{cell.Cell{0, 0, 1}}, board) != 1500 {
// 		t.Error("GetZoningScore should return 1500 for corner cell position")
// 	}
//
// 	if GetZoningScore([]cell.Cell{cell.Cell{0, 1, 1}}, board) != 1000 {
// 		t.Error("GetZoningScore should return 1500 for border cell position")
// 	}
//
// }
