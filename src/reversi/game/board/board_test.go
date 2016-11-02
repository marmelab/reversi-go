package board

import (
  "testing"
  "reflect"
  "reversi/game/cell"
)

func TestNewBoardShouldReturnNewBoard(t *testing.T) {

  board := New(4, 4)
  expectedBoard := Board{
    {0, 0, 0, 0},
    {0, 1, 2, 0},
    {0, 2, 1, 0},
    {0, 0, 0, 0},
  }

  if !reflect.DeepEqual(board, expectedBoard) {
    t.Error("New Board doesn't return expected Board")
  }

}

func TestIsFullShouldReturnTrueIfBoardIsFull(t *testing.T){

    board := New(1, 1)

    if(IsFull(board)){
      t.Error("IsFull Should'nt return true if Board is not filled")
    }

}

func TestIsFullShouldReturnFalseIfBoardIsNotFull(t *testing.T){

    board := New(1, 1)
    board[0][0] = cell.TypeBlack

    if(!IsFull(board)){
      t.Error("IsFull Should return true if Board is filled")
    }

}

func TestGetCellDistributionShouldReturnBoardCellDistribution(t *testing.T){

  board := New(8, 8)
  emptyExpectedDist := map[uint8]uint8{cell.TypeEmpty: uint8(60), cell.TypeBlack: uint8(2), cell.TypeWhite: uint8(2)}

  if !reflect.DeepEqual(GetCellDistribution(board), emptyExpectedDist){
    t.Error("Wrong Cell Distribution for new board")
  }

  board[0][0] = cell.TypeBlack
  board[1][0] = cell.TypeWhite
  board[2][0] = cell.TypeBlack
  filledExpectedDist := map[uint8]uint8{cell.TypeEmpty: uint8(57), cell.TypeBlack: uint8(4), cell.TypeWhite: uint8(3)}

  if !reflect.DeepEqual(GetCellDistribution(board), filledExpectedDist){
    t.Error("Wrong Cell Distribution")
  }

}

func TestGetDepartureCellsShouldReturnDepartureCells(t *testing.T){

  board := New(8, 8)
  expectedDepartureCells := []cell.Cell{
    cell.Cell{4, 4, cell.TypeBlack},
    cell.Cell{3, 3, cell.TypeBlack},
    cell.Cell{3, 4, cell.TypeWhite},
    cell.Cell{4, 3, cell.TypeWhite},
  }

  if !reflect.DeepEqual(GetDepartureCells(board), expectedDepartureCells){
    t.Error("expectedDepartureCells is not equal to the GetDepartureCells result")
  }

}

func TestGetSizeShouldReturnBoardSizeFromFirstRow(t *testing.T){

  board := Board{{0}, {0, 1, 3}}
  x, y := GetSize(board)

  if x != 1 || y != 2 {
    t.Error("The expected board size is not right")
  }

}
