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

func TestIsFinishedShouldReturnTrueIfThereIsNoOpportunityToPlay(t *testing.T) {

	blockedGame := Game{board.Board{{cell.TypeBlack, cell.TypeBlack, cell.TypeEmpty}}, make([]player.Player, 2, 2), 0}

	if !IsFinished(blockedGame) {
		t.Error("IsFinished should return true if there is no opportunity to play")
	}

}
