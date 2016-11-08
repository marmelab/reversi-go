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

	for y := 0; y < ySize; y++ {
		zoningScoreBoard = append(zoningScoreBoard, make([]int, xSize, xSize))
		for x := 0; x < xSize; x++ {

			zonScore = 0

			// Helpers

			isAroundCornerVertical := (x == 1 && (y < 2 || y > ySize-3)) || (x == xSize-2 && (y < 2 || y > ySize-3))
			isAroundCornerHorizontal := (y == 1 && (x < 2 || x > xSize-3)) || (y == ySize-2 && (x < 2 || x > xSize-3))
			isAroundCorner := isAroundCornerVertical || isAroundCornerHorizontal

			// Borders (except around corner)

			if (x == 0 || x == xSize-1 || y == 0 || y == ySize-1) && !isAroundCorner {
				zonScore += 50
			}

			// Center zoneScoringBoard

			if (x > 1 && x < xSize-2 && y == 2) || (x > 1 && x < xSize-2 && y == ySize-3) || (y > 1 && y < ySize-2 && x == xSize-3) || (y > 1 && y < ySize-2 && x == 2) {
				zonScore += 50
			}

			// Corner

			if (x == 0 && y == 0) || (x == xSize-1 && y == ySize-1) || (x == 0 && y == ySize-1) || (x == xSize-1 && y == 0) {
				zonScore += 150
			}

			// Negate around corners

			if isAroundCorner {
				zonScore -= 50
			}

			zoningScoreBoard[y][x] = zonScore

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
