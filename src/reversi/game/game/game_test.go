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

func TestIsFinishedShouldReturnSameAsBoardIsFull(t *testing.T) {

	fullBoard := board.Board{{cell.TypeBlack}}
	emptyBoard := board.Board{{cell.TypeEmpty}}

	finishedGame := Game{fullBoard, make([]player.Player, 2, 2), 0}
	notFinishedGame := Game{emptyBoard, make([]player.Player, 2, 2), 0}

	if IsFinished(finishedGame) != board.IsFull(fullBoard) || IsFinished(notFinishedGame) != board.IsFull(emptyBoard) {
		t.Error("IsFinished don't return same value as Board IsFull")
	}

}
