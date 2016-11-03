package ai

import (
	"reversi/game/cell"
	"reversi/game/game"
	"reversi/game/player"
)

func BestCellChange(currentGame game.Game, aiPlayer player.Player, depth int, depthLimit int, cellChangeChan chan cell.Cell) int {

	if game.IsFinished(currentGame) || depth >= depthLimit {
		return Score(currentGame, aiPlayer, depth)
	}

	depth++

	var scores []int
	var cellChanges []cell.Cell
	var possibleNewGame game.Game
	var resScore int
	var resScoreIdx int

	for _, availableCellChange := range game.GetAvailableCellChanges(currentGame) {
		possibleNewGame = game.PlayCellChange(currentGame, availableCellChange)
		possibleNewGame, _ = game.SwitchPlayer(possibleNewGame)
		scores = append(scores, BestCellChange(possibleNewGame, aiPlayer, depth, depthLimit, cellChangeChan))
		cellChanges = append(cellChanges, availableCellChange)
	}

	if game.GetCurrentPlayer(currentGame) == aiPlayer {
		resScore, resScoreIdx = MinIntSlice(scores)
	} else {
		resScore, resScoreIdx = MaxIntSlice(scores)
	}

	cellChangeChan <- cellChanges[resScoreIdx]

	return resScore

}

func Score(currentGame game.Game, aiPlayer player.Player, depth int) int {

	gameWinner, err := game.GetWinPlayer(currentGame)

	if aiPlayer == gameWinner {
		return 1000 - depth
	} else if err == nil {
		return -1000 + depth
	} else {
		return 0
	}

}

func MinIntSlice(v []int) (m int, idx int) {
	if len(v) > 0 {
		m = v[0]
		idx = 0
	}
	for i := 1; i < len(v); i++ {
		if v[i] < m {
			m = v[i]
			idx = i
		}
	}
	return
}

func MaxIntSlice(v []int) (m int, idx int) {
	if len(v) > 0 {
		m = v[0]
		idx = 0
	}
	for i := 1; i < len(v); i++ {
		if v[i] > m {
			m = v[i]
			idx = i
		}
	}
	return
}
