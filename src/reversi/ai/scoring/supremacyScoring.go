package scoring

import (
	"reversi/game/board"
	"reversi/game/cell"
)

func GetSupremacyScore(gameBoard board.Board, cellType uint8) int {

	cellDist := board.GetCellDistribution(gameBoard)
	reverseCellType := cell.GetReverseCellType(cellType)

	// Score based on the number of cell of the player's cellType
	// Nb of player cells - Nb of opponent cells - number of possibilities
	// -(boardX*boardY) < score < boardX*boardY

	return int(cellDist[cellType]) - int(cellDist[reverseCellType]) - int(cellDist[cell.TypeEmpty])

}
