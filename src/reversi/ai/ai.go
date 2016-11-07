package ai

import (
	"math"
	"reversi/game/board"
	"reversi/game/cell"
	"reversi/game/game"
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

	cellDist := board.GetCellDistribution(gameBoard)
	reverseCellType := cell.GetReverseCellType(aiPlayer.CellType)

	var winScore int

	if cellDist[aiPlayer.CellType] > cellDist[reverseCellType] {
		winScore = ScoringLevelLimit - depth
	} else if cellDist[aiPlayer.CellType] < cellDist[reverseCellType] {
		winScore = -ScoringLevelLimit + depth
	} else {
		winScore = 0
	}

	// Enhance with "techniques particulières à Othello"
	// http://www.ffothello.org/informatique/algorithmes/

	return winScore

}
