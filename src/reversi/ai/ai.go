package ai

import (
	"errors"
	"fmt"
	"math"
	"reversi/ai/scoring"
	"reversi/debug"
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
	Detail      map[string]int
}

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
	for i := 0; i < 5; i++ {
		go ScoringWorker(nodes, scores)
	}

	// Start board graph visitors
	for _, cellChange := range legalCellChanges {
		go RecursiveNodeVisitor(Node{currentBoard, cellChange, cellChange, false, cellType, 0}, nodes)
	}

	return CaptureBestCellChange(scores, timeout), nil

}

func ScoringWorker(nodes <-chan Node, scores chan<- Scoring) {
	for node := range nodes {
		start := time.Now()
		score, details := Score(node)
		scores <- Scoring{node, time.Since(start), score, details}
	}
}

func CaptureBestCellChange(scores chan Scoring, stopProcess chan bool) cell.Cell {

	aggregatedScores := map[cell.Cell]int{}
	finished := false

	for !finished {
		select {
		case finished = <-stopProcess:
		case scoring := <-scores:
			rcc := scoring.ScoreNode.RootCellChange
			debug.Log(fmt.Sprintf("%d:%d (from %d depth) - Score: %d (%s)", rcc.X+1, rcc.Y+1, scoring.ScoreNode.Depth, scoring.Score, debug.MapFormat(scoring.Detail)))
			if _, ok := aggregatedScores[rcc]; !ok {
				aggregatedScores[rcc] = 0
			}
			aggregatedScores[rcc] += scoring.Score
		}
	}

	bestCellChange := cell.Cell{}
	maxScore := -math.MaxInt32

	for cellChange, score := range aggregatedScores {
		debug.Log(fmt.Sprintf("## Aggregated %d:%d => %d", cellChange.X+1, cellChange.Y+1, score))
		if score >= maxScore {
			maxScore = score
			bestCellChange = cellChange
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

func Score(node Node) (int, map[string]int) {

	// Enhance with "techniques particulières à Othello"
	// http://www.ffothello.org/informatique/algorithmes/
	// http://www.ffothello.org/othello/principes-strategiques/

	availableCellChanges := board.GetLegalCellChangesForCellType(node.CellType, node.Board)

	zoningScore := scoring.GetZoningScore(availableCellChanges, node.Board, node.Depth)
	supremacyScore := scoring.GetSupremacyScore(node.Board, node.CellType)
	possibilitiesScore := scoring.GetPossibilitiesScore(len(availableCellChanges), node.Depth)

	totalScore := zoningScore + supremacyScore + possibilitiesScore

	details := map[string]int{
		"zoning":        zoningScore,
		"supremacy":     supremacyScore,
		"possibilities": possibilitiesScore,
	}

	if node.IsOpponent {
		return -totalScore, details
	}

	return totalScore, details

}

func GetBoardFromCellChange(currentBoard board.Board, cellChange cell.Cell) board.Board {
	cellChangesToApply := append(board.GetFlippedCellsFromCellChange(cellChange, currentBoard), cellChange)
	return board.ComputeCells(cellChangesToApply, currentBoard)
}
