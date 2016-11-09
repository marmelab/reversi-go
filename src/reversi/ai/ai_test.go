package ai

import (
	//"fmt"
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

func TestNodeVisitorShouldReturnValidNodeChan(t *testing.T) {

	currBoard, _ := board.InitCells(board.New(8, 8))
	nodes := NodeVisitor(Node{currBoard, cell.Cell{}, cell.Cell{}, false, cell.TypeWhite, 0})

	throwErr := func() {
		t.Error("NodeVisitor should return Node Chan with valid values")
	}

	if n := <-nodes; (n.cellChange != cell.Cell{3, 2, cell.TypeWhite}) {
		throwErr()
	}

	if n := <-nodes; (n.cellChange != cell.Cell{2, 3, cell.TypeWhite}) {
		throwErr()
	}

	if n := <-nodes; (n.cellChange != cell.Cell{5, 4, cell.TypeWhite}) {
		throwErr()
	}

	if n := <-nodes; (n.cellChange != cell.Cell{4, 5, cell.TypeWhite}) {
		throwErr()
	}

}

func TestRecursiveNodeVisitorShouldRecursivelyVisitNode(t *testing.T) {

	nodes := make(chan Node, 1)
	currBoard, _ := board.InitCells(board.New(8, 8))

	RecursiveNodeVisitor(Node{currBoard, cell.Cell{}, cell.Cell{}, false, cell.TypeWhite, 0}, nodes)

	countFirstLevel := 0
	countSecondLevel := 0
	countThirdLevel := 0

	for i := 0; i < 10000; i++ {

		node := <-nodes

		if node.Depth == 1 {
			countFirstLevel += 1
		}

		if node.Depth == 2 {
			countSecondLevel += 1
		}

		if node.Depth == 3 {
			countThirdLevel += 1
		}

	}

	if countFirstLevel != 4 || countSecondLevel != 12 || countThirdLevel != 56 {
		t.Error("RecursiveNodeVisitor should visit node recursively")
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
		{200, -50, 50, 50, 50, 50, -50, 200},
		{-50, -50, 0, 0, 0, 0, -50, -50},
		{50, 0, 50, 50, 50, 50, 0, 50},
		{50, 0, 50, 0, 0, 50, 0, 50},
		{50, 0, 50, 0, 0, 50, 0, 50},
		{50, 0, 50, 50, 50, 50, 0, 50},
		{-50, -50, 0, 0, 0, 0, -50, -50},
		{200, -50, 50, 50, 50, 50, -50, 200},
	}

	zoneScoringBoard := BuildZoneScoringBoard(8, 8)

	//fmt.Println(zoneScoringBoard)

	if !reflect.DeepEqual(zoneScoringBoard, expectedZoneScoringBoard) {
		t.Error("BuildZoneScoringBoard should return a valid score matrix")
	}

}
