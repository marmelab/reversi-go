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

func TestNodeVisitorShouldReturnValidNodeChan(t *testing.T) {

	currBoard, _ := board.InitCells(board.New(8, 8))
	nodes := NodeVisitor(Node{currBoard, cell.Cell{}, cell.Cell{}, false, cell.TypeWhite, 0})

	throwErr := func() {
		t.Error("NodeVisitor should return Node Chan with valid values")
	}

	if len(nodes) != 4 {
		throwErr()
	}

	if n := nodes[0]; (n.cellChange != cell.Cell{3, 2, cell.TypeWhite}) {
		throwErr()
	}

	if n := nodes[1]; (n.cellChange != cell.Cell{2, 3, cell.TypeWhite}) {
		throwErr()
	}

	if n := nodes[2]; (n.cellChange != cell.Cell{5, 4, cell.TypeWhite}) {
		throwErr()
	}

	if n := nodes[3]; (n.cellChange != cell.Cell{4, 5, cell.TypeWhite}) {
		throwErr()
	}

}

func TestRecursiveNodeVisitorShouldRecursivelyVisitNode(t *testing.T) {

	nodes := make(chan Node, 1)
	currBoard, _ := board.InitCells(board.New(8, 8))

	go RecursiveNodeVisitor(Node{currBoard, cell.Cell{}, cell.Cell{}, false, cell.TypeWhite, 0}, nodes)

	countFirstLevel := 0
	countSecondLevel := 0

	for i := 0; i < 10000; i++ {

		node := <-nodes

		if node.Depth == 1 {
			countFirstLevel += 1
		}

		if node.Depth == 2 {
			countSecondLevel += 1
		}

	}

	if countFirstLevel != 4 || countSecondLevel != 12{
		t.Error("RecursiveNodeVisitor should visit node recursively")
	}

}
