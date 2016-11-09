package game

import (
	"reflect"
	"reversi/game/board"
	"reversi/game/cell"
	"reversi/game/player"
	"testing"
)

func TestNewGameShouldReturnNewGame(t *testing.T) {

	game := New(make([]player.Player, 2, 2))
	expectedBoard, _ := board.InitCells(board.New(8, 8))
	expectedGame := Game{expectedBoard, make([]player.Player, 2, 2), 0}

	if !reflect.DeepEqual(game, expectedGame) {
		t.Error("New doesn't return expected game struct")
	}

}

func TestIsFinishedShouldReturnTrueIfBoardIsFull(t *testing.T) {

	fullBoardGame := Game{board.Board{{cell.TypeBlack}}, make([]player.Player, 2, 2), 0}

	if !IsFinished(fullBoardGame) {
		t.Error("IsFinished should return true if board is full")
	}

}

func TestIsFinishedShouldReturnFalseIfBoardIsNotFull(t *testing.T) {

	notFullBoardGame := Game{board.Board{{cell.TypeEmpty}}, make([]player.Player, 2, 2), 0}

	if IsFinished(notFullBoardGame) {
		t.Error("IsFinished should return false if board is not full")
	}

}
