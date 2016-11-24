package scoring

import (
	"reversi/game/board"
	"reversi/game/cell"
	"reversi/game/matrix"
)

func GetZoningScore(availableCellChanges []cell.Cell, gameBoard board.Board) int {

	zoningScore := 0
	xSize, ySize := matrix.GetSize(gameBoard)

	//Generate zoning score board

	zoningScoreBoard := BuildZoneScoringBoard(xSize, ySize)

	for _, cellChange := range availableCellChanges {

		xPos := int(cellChange.X)
		yPos := int(cellChange.Y)

		zoningScore += zoningScoreBoard[yPos][xPos]

	}

	return int(zoningScore)

}

func BuildZoneScoringBoard(xSize int, ySize int) [][]int {

	// ----------------------------------------------
	// | 250 | -50 | 50 | 50 | 50 | 50 | -50 | 250 |
	// ----------------------------------------------
	// | -50 | -50 | 0  | 0  | 0  | 0  | -50 | -50 |
	// ----------------------------------------------
	// | 50  |  0  | 50 | 50 | 50 | 50 |  0  | 50  |
	// ----------------------------------------------
	// | 50  |  0  | 50 |  0 | 0  | 50 |  0  | 50  |
	// ----------------------------------------------
	// | 50  |  0  | 50 |  0 | 0  | 50 |  0  | 50  |
	// ----------------------------------------------
	// | 50  |  0  | 50 | 50 | 50 | 50 |  0  | 50  |
	// ----------------------------------------------
	// | -50 | -50 | 0  | 0  | 0  | 0  | -50 | -50 |
	// ----------------------------------------------
	// | 250 | -50 | 50 | 50 | 50 | 50 | -50 | 250 |
	// ----------------------------------------------

	zoningScoreBoard := [][]int{}
	var zonScore int

	for y := 0; y < ySize; y++ {
		zoningScoreBoard = append(zoningScoreBoard, make([]int, xSize, xSize))
		for x := 0; x < xSize; x++ {

			zonScore = 0

			// Borders (except around corner)
			if isOnBorder(x, y, xSize, ySize) && !isAroundCorner(x, y, xSize, ySize) {
				zonScore += 50
			}

			// Center zoneScoringBoard
			if (x > 1 && x < xSize-2 && y == 2) || (x > 1 && x < xSize-2 && y == ySize-3) || (y > 1 && y < ySize-2 && x == xSize-3) || (y > 1 && y < ySize-2 && x == 2) {
				zonScore += 50
			}

			// Corner
			if isCorner(x, y, xSize, ySize) {
				zonScore += 250
			}

			// Negate around corners
			if isAroundCorner(x, y, xSize, ySize) {
				zonScore -= 50
			}

			zoningScoreBoard[y][x] = zonScore

		}
	}

	return zoningScoreBoard

}

func isCorner(x int, y int, xSize int, ySize int) bool {
	return (x == 0 && y == 0) || (x == xSize-1 && y == ySize-1) || (x == 0 && y == ySize-1) || (x == xSize-1 && y == 0)
}

func isAroundCorner(x int, y int, xSize int, ySize int) bool {
	isAroundCornerVertical := (x == 1 && (y < 2 || y > ySize-3)) || (x == xSize-2 && (y < 2 || y > ySize-3))
	isAroundCornerHorizontal := (y == 1 && (x < 2 || x > xSize-3)) || (y == ySize-2 && (x < 2 || x > xSize-3))
	return isAroundCornerVertical || isAroundCornerHorizontal
}

func isOnBorder(x int, y int, xSize int, ySize int) bool {
	return x == 0 || x == xSize-1 || y == 0 || y == ySize-1
}
