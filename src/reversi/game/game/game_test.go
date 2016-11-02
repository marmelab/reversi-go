package game

import (
  "testing"
  "reflect"
  "reversi/game/player"
  "reversi/game/board"
  "reversi/game/cell"
)

func TestNewGameShouldReturnNewGame(t *testing.T) {

  game := New(1, 1, make([]player.Player, 2, 2))
  expectedGame := Game{board.New(1, 1),  make([]player.Player, 2, 2), 0}

  if !reflect.DeepEqual(game, expectedGame){
    t.Error("New doesn't return expected game struct")
  }

}

func TestIsFinishedShouldReturnSameAsBoardIsFull(t *testing.T){

  fullBoard := board.Board{{cell.TypeBlack}}
  emptyBoard := board.Board{{cell.TypeEmpty}}

  finishedGame := Game{fullBoard,  make([]player.Player, 2, 2), 0}
  notFinishedGame := Game{emptyBoard,  make([]player.Player, 2, 2), 0}

  if IsFinished(finishedGame) != board.IsFull(fullBoard) || IsFinished(notFinishedGame) != board.IsFull(emptyBoard){
    t.Error("IsFinished don't return same value as Board IsFull")
  }

}
