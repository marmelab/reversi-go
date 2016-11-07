package ai

import (
	"math"
	"reversi/game/board"
	"reversi/game/cell"
	"reversi/game/game"
	"reversi/game/matrix"
	"reversi/game/player"
)

const ScoringLevelLimit int = math.MaxInt8

func GetMaxScore(currentGame game.Game, aiPlayer player.Player, depth int, depthLimit int) int {

	if game.IsFinished(currentGame) || depth >= depthLimit {
		return Score(currentGame.Board, aiPlayer, depth)
	}

	reversePlayerMaxScore := -ScoringLevelLimit

	for _, cellChange := range game.GetAvailableCellChanges(currentGame) {

		virtualGame, _ := game.PlayTurn(currentGame, cellChange)
		reversePlayerScore := GetMinScore(virtualGame, aiPlayer, depth+1, depthLimit)

		if reversePlayerScore > reversePlayerMaxScore {
			reversePlayerMaxScore = reversePlayerScore
		}

	}

	return reversePlayerMaxScore

}

func GetMinScore(currentGame game.Game, aiPlayer player.Player, depth int, depthLimit int) int {

	if game.IsFinished(currentGame) || depth >= depthLimit {
		return Score(currentGame.Board, aiPlayer, depth)
	}

	reversePlayerMinScore := ScoringLevelLimit

	for _, cellChange := range game.GetAvailableCellChanges(currentGame) {

		virtualGame, _ := game.PlayTurn(currentGame, cellChange)
		reversePlayerScore := GetMaxScore(virtualGame, aiPlayer, depth+1, depthLimit)

		if reversePlayerScore < reversePlayerMinScore {
			reversePlayerMinScore = reversePlayerScore
		}

	}

	return reversePlayerMinScore

}

func GetBestCellChange(currentGame game.Game, aiPlayer player.Player, depth int, depthLimit int) cell.Cell {

	maxScore := -ScoringLevelLimit
	bestCellChange := cell.Cell{}

	for _, cellChange := range game.GetAvailableCellChanges(currentGame) {

		virtualGame, _ := game.PlayTurn(currentGame, cellChange)
		cellChangeScore := GetMaxScore(virtualGame, aiPlayer, depth, depthLimit)

		if cellChangeScore > maxScore {
			maxScore = cellChangeScore
			bestCellChange = cellChange
		}

	}

	return bestCellChange

}

func Score(gameBoard board.Board, aiPlayer player.Player, depth int) int {

	// Enhance with "techniques particulières à Othello"
	// http://www.ffothello.org/informatique/algorithmes/

	availableCellChanges := board.GetLegalCellChangesForCellType(aiPlayer.CellType, gameBoard)

	supremacyScore := GetSupremacyScore(gameBoard, aiPlayer.CellType, depth)
	zoningScore := GetZoningScore(availableCellChanges, gameBoard)

	return supremacyScore + zoningScore

}

func GetZoningScore(availableCellChanges []cell.Cell, gameBoard board.Board) int {

	zoningScore := 0
	xSize, ySize := matrix.GetSize(gameBoard)

	// Scoring Strategy
	// +1000 for board limits
	// +1500 for board corners

	for _, cellChange := range availableCellChanges {
		xPos := int(cellChange.X)
		yPos := int(cellChange.Y)
		if xPos == 0 || xPos == xSize-1 || yPos == 0 || yPos == ySize-1 {
			zoningScore += 1000
			if (yPos == 0 && xPos == 0) || (yPos == ySize-1 && xPos == xSize-1) || (yPos == ySize-1 && xPos == 0) || (yPos == 0 && xPos == xSize-1) {
				zoningScore += 500
			}
		}
	}

	return zoningScore

}

func GetSupremacyScore(gameBoard board.Board, cellType uint8, depth int) int {

	cellDist := board.GetCellDistribution(gameBoard)
	reverseCellType := cell.GetReverseCellType(cellType)

	// Score based on the number of cell of the player's cellType
	// The depth parameter permit to highlight near distribution configurations

	if cellDist[cellType] > cellDist[reverseCellType] {
		return ScoringLevelLimit - depth
	}

	if cellDist[cellType] < cellDist[reverseCellType] {
		return -ScoringLevelLimit + depth
	}

	return 0

}
