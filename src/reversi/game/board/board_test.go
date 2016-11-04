package board

import (
	"reflect"
	"reversi/game/cell"
	"testing"
)

func TestNewBoardShouldReturnNewBoard(t *testing.T) {

	board := New(4, 4)
	expectedBoard := Board{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	if !reflect.DeepEqual(board, expectedBoard) {
		t.Error("New Board doesn't return expected Board")
	}

}

func TestInitCellsShouldReturnErrorWithAnInvalidBoardSize(t *testing.T) {

	board := Board{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	_, err := InitCells(board)

	if err == nil {
		t.Error("InitCell should return an error with an invalid board size")
	}

}

func TestInitCellsShouldDrawDepartureCellsWithAValidBoardSize(t *testing.T) {

	board := Board{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	expectedBoard := Board{
		{0, 0, 0, 0},
		{0, 1, 2, 0},
		{0, 2, 1, 0},
		{0, 0, 0, 0},
	}

	newboard, _ := InitCells(board)

	if !reflect.DeepEqual(newboard, expectedBoard) {
		t.Error("InitCell should draw departure cells with valid board size")
	}

}

func TestIsFullShouldReturnTrueIfBoardIsFilled(t *testing.T) {

	board := New(1, 1)
	board[0][0] = cell.TypeBlack

	if !IsFull(board) {
		t.Error("IsFull should return true if board is filled")
	}

}

func TestIsFullShouldReturnFalseIfBoardIsNotFilled(t *testing.T) {

	board := New(1, 1)

	if IsFull(board) {
		t.Error("IsFull Should return false if Board is not filled")
	}

}

func TestGetCellDistributionShouldReturnBoardCellDistribution(t *testing.T) {

	board := New(8, 8)
	emptyExpectedDist := map[uint8]uint8{cell.TypeEmpty: uint8(64), cell.TypeBlack: uint8(0), cell.TypeWhite: uint8(0)}

	if !reflect.DeepEqual(GetCellDistribution(board), emptyExpectedDist) {
		t.Error("Wrong Cell Distribution for new board")
	}

	board[0][0] = cell.TypeBlack
	board[1][0] = cell.TypeWhite
	board[2][0] = cell.TypeBlack
	filledExpectedDist := map[uint8]uint8{cell.TypeEmpty: uint8(61), cell.TypeBlack: uint8(2), cell.TypeWhite: uint8(1)}

	if !reflect.DeepEqual(GetCellDistribution(board), filledExpectedDist) {
		t.Error("Wrong Cell Distribution")
	}

}

func TestGetDepartureCellsShouldReturnDepartureCells(t *testing.T) {

	board := New(8, 8)
	expectedDepartureCells := []cell.Cell{
		cell.Cell{4, 4, cell.TypeBlack},
		cell.Cell{3, 3, cell.TypeBlack},
		cell.Cell{3, 4, cell.TypeWhite},
		cell.Cell{4, 3, cell.TypeWhite},
	}

	if !reflect.DeepEqual(GetDepartureCells(board), expectedDepartureCells) {
		t.Error("expectedDepartureCells is not equal to the GetDepartureCells result")
	}

}

func DrawCellsShouldDrawCellOnBoard(t *testing.T) {

	expectedBoard := Board{{cell.TypeBlack}}
	board := DrawCells([]cell.Cell{cell.New(uint8(0), uint8(0), cell.TypeBlack)}, Board{})

	if !reflect.DeepEqual(board, expectedBoard) {
		t.Error("DrawCells should draw cells and return a new board")
	}

}

func TestGetLegalCellChangesForCellTypeShouldReturnLegalCellChangeSlice(t *testing.T) {

	board := Board{{cell.TypeBlack, cell.TypeWhite, cell.TypeEmpty}}
	expectedCellChanges := []cell.Cell{cell.Cell{2, 0, cell.TypeBlack}}

	if !reflect.DeepEqual(GetLegalCellChangesForCellType(cell.TypeBlack, board), expectedCellChanges) {
		t.Error("GetLegalCellChanges sould return the expected Cell Changes")
	}

}

func TestGetCellTypeShouldReturnCellType(t *testing.T) {

	board := Board{{cell.TypeBlack, cell.TypeWhite, cell.TypeEmpty}}

	if GetCellType(0, 0, board) != cell.TypeBlack || GetCellType(1, 0, board) != cell.TypeWhite {
		t.Error("GetCellType should return CellType")
	}

}

func TestGetCellTypeShouldReturnEmptyCellTypeForOutOfRangeCell(t *testing.T) {

	board := Board{{cell.TypeBlack, cell.TypeWhite, cell.TypeEmpty}}

	if GetCellType(42, 42, board) != cell.TypeEmpty {
		t.Error("GetCellType should return EmptyCellType for OutOfRange cell")
	}

}
