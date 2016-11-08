package ai

import (
	"errors"
	"math"
	"reversi/game/board"
	"reversi/game/cell"
	"reversi/game/matrix"
	"time"
)

type Node struct {
	Board          board.Board
	RootCellChange cell.Cell
	IsOpponent     bool
	CellType       uint8
	Depth          int
}

func GetBestCellChangeInTime(currentBoard board.Board, cellType uint8, duration time.Duration) (cell.Cell, error) {

	nodes := make(chan Node)
	timeout := make(chan bool, 1)
	bestCellChange := cell.Cell{}

	go func() {
		time.Sleep(duration)
		timeout <- true
	}()

	legalCellChanges := board.GetLegalCellChangesForCellType(cellType, currentBoard)

	if len(legalCellChanges) == 0 {
		return bestCellChange, errors.New("There's no legal cell change for this cellType.")
	}

	for _, cellChange := range legalCellChanges {
		RecursiveNodeVisitor(Node{currentBoard, cellChange, false, cellType, 1}, nodes)
	}

	finished := false
	maxScore := -math.MaxInt32

	for !finished {
		select {
		case finished = <-timeout:
		case node := <-nodes:
			score := Score(node)
			if score > maxScore {
				maxScore = score
				bestCellChange = node.RootCellChange
			}
		}
	}

	return bestCellChange, nil

}

func NodeVisitor(node Node) chan Node {
	out := make(chan Node)
	go func() {
		legalCellChanges := board.GetLegalCellChangesForCellType(node.CellType, node.Board)
		for _, cellChange := range legalCellChanges {
			nodeBoard := GetBoardFromCellChange(node.Board, cellChange)
			out <- Node{nodeBoard, node.RootCellChange, !node.IsOpponent, cell.GetReverseCellType(node.CellType), node.Depth + 1}
		}
		close(out)
	}()
	return out
}

func RecursiveNodeVisitor(node Node, out chan Node) {
	go func() {
		visitorChannel := NodeVisitor(node)
		for visitedNode := range visitorChannel {
			out <- visitedNode
			RecursiveNodeVisitor(visitedNode, out)
		}
	}()
}

func Score(node Node) int {

	// Enhance with "techniques particulières à Othello"
	// http://www.ffothello.org/informatique/algorithmes/

	availableCellChanges := board.GetLegalCellChangesForCellType(node.CellType, node.Board)

	zoningScore := GetZoningScore(availableCellChanges, node.Board)
	supremacyScore := GetSupremacyScore(node.Board, node.CellType)

	return zoningScore + supremacyScore

}

func GetBoardFromCellChange(currentBoard board.Board, cellChange cell.Cell) board.Board {
	cellChangesToApply := append(board.GetFlippedCellsFromCellChange(cellChange, currentBoard), cellChange)
	return board.DrawCells(cellChangesToApply, currentBoard)
}

func GetZoningScore(availableCellChanges []cell.Cell, gameBoard board.Board) int {

	zoningScore := 0
	xSize, ySize := matrix.GetSize(gameBoard)

	//Generate zoning score board

	zoningScoreBoard := BuildZoneScoringBoard(xSize, ySize)

	// Scoring Strategy
	// +50 for board limits (except around corners)
	// +100 for board corners

	for _, cellChange := range availableCellChanges {

		xPos := int(cellChange.X)
		yPos := int(cellChange.Y)

		zoningScore += zoningScoreBoard[yPos][xPos]

	}

	return zoningScore

}

func BuildZoneScoringBoard(xSize int, ySize int) [][]int {

	zoningScoreBoard := [][]int{}
	var zonScore int

	for zonY := 0; zonY < ySize; zonY++ {
		zoningScoreBoard = append(zoningScoreBoard, make([]int, xSize, xSize))
		for zonX := 0; zonX < xSize; zonX++ {

			zonScore = 0

			// Borders (except around corners)
			if zonX == 0 || zonX == xSize-1 || zonY == 0 || zonY == ySize-1 {
				zonScore += 50
			}

			// Center zoneScoringBoard
			if (zonX > 1 && zonX < xSize-2 && zonY == 2) || (zonX > 1 && zonX < xSize-2 && zonY == ySize-3) || (zonY > 1 && zonY < ySize-2 && zonX == xSize-3) || (zonY > 1 && zonY < ySize-2 && zonX == 2) {
				zonScore += 50
			}

			// Corner
			if (zonX == 0 && zonY == 0) || (zonX == xSize-1 && zonY == ySize-1) || (zonX == 0 && zonY == ySize-1) || (zonX == xSize-1 && zonY == 0) {
				zonScore += 150
			}

			zoningScoreBoard[zonY][zonX] = zonScore

		}
	}

	return zoningScoreBoard

}

func GetSupremacyScore(gameBoard board.Board, cellType uint8) int {

	cellDist := board.GetCellDistribution(gameBoard)
	reverseCellType := cell.GetReverseCellType(cellType)

	// Score based on the number of cell of the player's cellType
	// Nb of player cells - Nb of opponent cells - number of possibilities
	// -(boardX*boardY) < score < boardX*boardY

	return int(cellDist[cellType]) - int(cellDist[reverseCellType]) - int(cellDist[cell.TypeEmpty])

}
