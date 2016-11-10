package ai

import (
	"errors"
	"reversi/ai/scoring"
	"reversi/game/board"
	"reversi/game/cell"
	"time"
)

type Node struct {
	Board          board.Board
	cellChange     cell.Cell
	RootCellChange cell.Cell
	IsOpponent     bool
	CellType       uint8
	Depth          int
}

type Scoring struct {
	ScoreNode   Node
	ScoringTime time.Duration
	Score       int
}

const SCORING_WORKER_COUNT int = 4

func GetBestCellChangeInTime(currentBoard board.Board, cellType uint8, duration time.Duration) (cell.Cell, error) {

	nodes := make(chan Node, 100)
	scores := make(chan Scoring)
	timeout := make(chan bool, 1)

	go func() {
		time.Sleep(duration)
		timeout <- true
	}()

	legalCellChanges := board.GetLegalCellChangesForCellType(cellType, currentBoard)

	if len(legalCellChanges) == 0 {
		return cell.Cell{}, errors.New("There's no legal cell change for this cellType.")
	}

	if len(legalCellChanges) == 1 {
		return legalCellChanges[0], nil
	}

	// Start scoring workers

	for i := 0; i < SCORING_WORKER_COUNT; i++ {
		go ScoringWorker(nodes, scores)
	}

	// Start board graph visitors

	for _, cellChange := range legalCellChanges {
		go RecursiveNodeVisitor(Node{currentBoard, cellChange, cellChange, false, cellType, 1}, nodes)
	}

	return CaptureBestCellChange(scores, timeout), nil

}

func ScoringWorker(nodes <-chan Node, scores chan<- Scoring) {
	for node := range nodes {
		start := time.Now()
		score := Score(node)
		scores <- Scoring{node, time.Since(start), score}
	}
}

func CaptureBestCellChange(scores chan Scoring, stopProcess chan bool) cell.Cell {

	bestCellChange := cell.Cell{}
	finished := false
	maxScore := 0

	for !finished {
		select {
		case finished = <-stopProcess:
		case scoring := <-scores:
			if scoring.Score > maxScore {
				maxScore = scoring.Score
				bestCellChange = scoring.ScoreNode.RootCellChange
			}
		}
	}

	return bestCellChange

}

func RecursiveNodeVisitor(rootNode Node, out chan Node) {
	for _, node := range NodeVisitor(rootNode) {
		out <- node
		go RecursiveNodeVisitor(node, out)
	}
}

func NodeVisitor(node Node) []Node {
	out := []Node{}
	legalCellChanges := board.GetLegalCellChangesForCellType(node.CellType, node.Board)
	for _, cellChange := range legalCellChanges {
		nodeBoard := GetBoardFromCellChange(node.Board, cellChange)
		out = append(out, Node{nodeBoard, cellChange, node.RootCellChange, !node.IsOpponent, cell.GetReverseCellType(node.CellType), node.Depth + 1})
	}
	return out
}

func Score(node Node) int {

	// Enhance with "techniques particulières à Othello"
	// http://www.ffothello.org/informatique/algorithmes/

	availableCellChanges := board.GetLegalCellChangesForCellType(node.CellType, node.Board)

	zoningScore := scoring.GetZoningScore(availableCellChanges, node.Board)
	supremacyScore := scoring.GetSupremacyScore(node.Board, node.CellType)

	totalScore := zoningScore + supremacyScore

	if node.IsOpponent {
		return -totalScore
	}

	return totalScore

}

func GetBoardFromCellChange(currentBoard board.Board, cellChange cell.Cell) board.Board {
	cellChangesToApply := append(board.GetFlippedCellsFromCellChange(cellChange, currentBoard), cellChange)
	return board.ComputeCells(cellChangesToApply, currentBoard)
}
