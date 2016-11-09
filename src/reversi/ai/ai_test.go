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

func BenchmarkScore(b *testing.B) {

	currentBoard, _ := board.InitCells(board.New(8, 8))

	currentBoard[4][5] = cell.TypeBlack
	currentBoard[4][2] = cell.TypeWhite

	node := Node{currentBoard, cell.Cell{}, cell.Cell{}, false, cell.TypeWhite, 0}

	for n := 0; n < b.N; n++ {
		Score(node)
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

	nodes := make(chan Node, 100)
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

	if countFirstLevel != 4 || countSecondLevel != 12 {
		t.Error("RecursiveNodeVisitor should visit node recursively")
	}

}

func TestCaptureBestCellChangeShouldReturnCellChangeFromTheMaxScoredScore(t *testing.T) {

	scores := make(chan Scoring, 4)
	finish := make(chan bool, 1)

	scores <- Scoring{Node{board.Board{}, cell.Cell{}, cell.Cell{}, false, cell.TypeBlack, 1}, time.Second, 500}
	scores <- Scoring{Node{board.Board{}, cell.Cell{}, cell.Cell{0, 0, cell.TypeBlack}, false, cell.TypeBlack, 2}, time.Second, 6000}
	scores <- Scoring{Node{board.Board{}, cell.Cell{}, cell.Cell{}, false, cell.TypeBlack, 3}, time.Second, 2000}
	scores <- Scoring{Node{board.Board{}, cell.Cell{}, cell.Cell{}, false, cell.TypeBlack, 4}, time.Second, 1000}

	go func() {
		time.Sleep(time.Second)
		finish <- true
	}()

	if CaptureBestCellChange(scores, finish) != (cell.Cell{0, 0, cell.TypeBlack}) {
		t.Error("CaptureBestCellChange should capture best cell change from scores")
	}

}
