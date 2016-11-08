package ai

import (
	"fmt"
	"reflect"
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

func TestGetZoningScore(t *testing.T) {

	board, _ := board.InitCells(board.New(8, 8))

	if GetZoningScore([]cell.Cell{cell.Cell{0, 0, 1}}, board) != 200 {
		t.Error("GetZoningScore should return 200 for corner cell position")
	}

	if GetZoningScore([]cell.Cell{cell.Cell{0, 2, 1}}, board) != 50 {
		t.Error("GetZoningScore should return 50 for border cell position")
	}

}

func TestGetSupremacyScoreShouldReturnAValidSupremacyScore(t *testing.T) {

	if GetSupremacyScore(board.Board{{2, 2, 2, 2, 2}}, cell.TypeWhite) != 5 {
		t.Error("GetSupremacyScore should return valid score")
	}

	if GetSupremacyScore(board.Board{{1, 1, 1, 2, 2}}, cell.TypeWhite) != -1 {
		t.Error("GetSupremacyScore should return valid score")
	}

	if GetSupremacyScore(board.Board{{1, 1, 1, 2, 2, 0, 0, 0}}, cell.TypeWhite) != -4 {
		t.Error("GetSupremacyScore should return valid score")
	}

}

func TestBuildZoneScoringBoardShouldReturnAValidScoreMatrix(t *testing.T) {

	expectedZoneScoringBoard := [][]int{
		{200, 50, 50, 50, 50, 50, 50, 200},
		{50, 0, 0, 0, 0, 0, 0, 50},
		{50, 0, 50, 50, 50, 50, 0, 50},
		{50, 0, 50, 0, 0, 50, 0, 50},
		{50, 0, 50, 0, 0, 50, 0, 50},
		{50, 0, 50, 50, 50, 50, 0, 50},
		{50, 0, 0, 0, 0, 0, 0, 50},
		{200, 50, 50, 50, 50, 50, 50, 200},
	}

	zoneScoringBoard := BuildZoneScoringBoard(8, 8)

	if !reflect.DeepEqual(zoneScoringBoard, expectedZoneScoringBoard) {
		t.Error("BuildZoneScoringBoard should return a valid score matrix")
	}

}
