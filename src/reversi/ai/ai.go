package ai

import (
	"errors"
	"math"
	"reversi/ai/scoring"
	"reversi/ai/node"
	//"reversi/debug"
	"reversi/game/board"
	"reversi/game/cell"
	"github.com/hishboy/gocommons/lang"
	"time"
)

func GetBestCellChangeInTime(currentBoard board.Board, cellType uint8, duration time.Duration) (cell.Cell, error) {

	timeout := make(chan bool, 1)

	go func() {
		time.Sleep(duration)
		timeout <- true
	}()

	legalCellChanges := board.GetLegalCellChangesForCellType(cellType, currentBoard)

	// Handle special cases (1 or 0 possibility)

	if len(legalCellChanges) == 0 {
		return cell.Cell{}, errors.New("There's no legal cell change for this cellType.")
	}

	if len(legalCellChanges) == 1 {
		return legalCellChanges[0], nil
	}

	// Explore node tree for each root node

	nodes := make([]*node.Node, len(legalCellChanges))

	for i, cellChange := range legalCellChanges {
		rootNode := node.New(false, currentBoard, cellChange, cellChange.CellType)
		nodes[i] = &rootNode
		VisitNode(&rootNode, timeout)
	}

	return CaptureBestCellChange(nodes, timeout), nil

}

func CaptureBestCellChange(nodes []*node.Node, stopProcess chan bool) cell.Cell {

	finished := false

	for !finished {
		select {
		case finished = <-stopProcess:
		}
	}

	bestCellChange := cell.Cell{}
	maxScore := -math.MaxInt32

	for _, currNode := range nodes {
		currNode.Evaluate()
		if currNode.Score != nil && *currNode.Score >= maxScore {
			maxScore = *currNode.Score
			bestCellChange = currNode.GetRootNode().CellChange
		}
	}

	return bestCellChange

}

func VisitNode(rootNode *node.Node, stopProcess chan bool) {

	visitQueue := lang.NewQueue()
	visitQueue.Push(rootNode)

	go BfsNodeVisitor(visitQueue, stopProcess)

}

func BfsNodeVisitor(visitQueue *lang.Queue, stopProcess chan bool) {

	for visitQueue.Len() > 0 {

		currNode := visitQueue.Poll().(*node.Node)
		legalCellChanges := board.GetLegalCellChangesForCellType(currNode.CellType, currNode.Board)

		for _, cellChange := range legalCellChanges {
			select {
			    case <- stopProcess:
			        return
			    default:
					childNode := currNode.Add(GetBoardFromCellChange(currNode.Board, cellChange), cellChange)
					score, _ := Score(childNode)
					childNode.Score = score
					visitQueue.Push(childNode)
			}
		}

	}

}

func Score(currNode *node.Node) (*int, map[string]int) {

	// Enhance with "techniques particulières à Othello"
	// http://www.ffothello.org/informatique/algorithmes/
	// http://www.ffothello.org/othello/principes-strategiques/

	availableCellChanges := board.GetLegalCellChangesForCellType(currNode.CellType, currNode.Board)

	zoningScore := scoring.GetZoningScore(availableCellChanges, currNode.Board)
	supremacyScore := scoring.GetSupremacyScore(currNode.Board, currNode.CellType)
	possibilitiesScore := scoring.GetPossibilitiesScore(len(availableCellChanges))

	totalScore := zoningScore + supremacyScore + possibilitiesScore

	details := map[string]int{
		"zoning":        zoningScore,
		"supremacy":     supremacyScore,
		"possibilities": possibilitiesScore,
	}

	return &totalScore, details

}

func GetBoardFromCellChange(currentBoard board.Board, cellChange cell.Cell) board.Board {
	cellChangesToApply := append(board.GetFlippedCellsFromCellChange(cellChange, currentBoard), cellChange)
	return board.ComputeCells(cellChangesToApply, currentBoard)
}
