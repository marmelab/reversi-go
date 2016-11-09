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

	return zoningScore

}

func BuildZoneScoringBoard(xSize int, ySize int) [][]int {

	// ----------------------------------------------
	// | 200 | -50 | 50 | 50 | 50 | 50 | -50 | 200 |
	// ----------------------------------------------
	// | -50 | -50 | 0  | 0  | 0  | 0  | -50 | -50 |
	// ----------------------------------------------
	// | 50  |  0  | 50 | 50 | 50 | 50 | 0   | 50  |
	// ----------------------------------------------
	// | 50  |  0  | 50 |  0 | 0  | 50 |  0  | 50  |
	// ----------------------------------------------
	// | 50  |  0  | 50 |  0 | 0  | 50 |  0  | 50  |
	// ----------------------------------------------
	// | 50  |  0  | 50 | 50 | 50 | 50 |  0  | 50  |
	// ----------------------------------------------
	// | -50 | -50 | 0  | 0  | 0  | 0  | -50 | -50 |
	// ----------------------------------------------
	// | 200 | -50 | 50 | 50 | 50 | 50 | -50 | 200 |
	// ----------------------------------------------

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
