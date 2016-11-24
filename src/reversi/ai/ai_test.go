package ai

import (
	"reversi/game/board"
	"reversi/game/cell"
	"testing"
	"time"
)

func BenchmarkGetBestCellChange(b *testing.B) {

	currentBoard, _ := board.InitCells(board.New(8, 8))

	for n := 0; n < b.N; n++ {
		GetBestCellChangeInTime(currentBoard, cell.TypeBlack, time.Second)
	}

}

func TestGetBestCellChangeInTimeShouldReturnAnErrorIfThereAreNoPossibilities(t *testing.T) {

	board := board.Board{{cell.TypeEmpty, cell.TypeBlack, cell.TypeBlack, cell.TypeWhite}}
	_, err := GetBestCellChangeInTime(board, cell.TypeBlack, time.Second)

	if err == nil {
		t.Error("GetBestCellChangeInTime should return an error if there's no possibilities to play")
	}

}

func TestGetBestCellChangeInTimeShouldNotReturnAnErrorIfThereArePossibilities(t *testing.T) {

	_, err := board.InitCells(board.New(8, 8))

	if err != nil {
		t.Error("GetBestCellChangeInTime should not return an error if there are possibilities to play")
	}

}
