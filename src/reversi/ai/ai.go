package ai

import (
	"math"
	//"fmt"
	"reversi/game/board"
	"reversi/game/cell"
	"reversi/game/game"
	"reversi/game/player"
)

func GetMaxScore(currentGame game.Game, aiPlayer player.Player, depth int, depthLimit int) int {

	if game.IsFinished(currentGame) || depth >= depthLimit {
		return Score(currentGame.Board, aiPlayer, depth)
	}
	//fmt.Println("Max, Depth ", depth, " number of changes ", len( game.GetAvailableCellChanges(currentGame)))
	maxScore := -math.MaxInt32

	for _, cellChange := range game.GetAvailableCellChanges(currentGame) {

		virtualGame, _ := game.PlayTurn(currentGame, cellChange)
		reversePlayerScore := GetMinScore(virtualGame, aiPlayer, depth+1, depthLimit)

		if reversePlayerScore > maxScore {
			maxScore = reversePlayerScore
		}

	}

	return maxScore

}

func GetMinScore(currentGame game.Game, aiPlayer player.Player, depth int, depthLimit int) int {

	if game.IsFinished(currentGame) || depth >= depthLimit {
		return Score(currentGame.Board, aiPlayer, depth)
	}
	//fmt.Println("Min, Depth ", depth, " number of changes ", len( game.GetAvailableCellChanges(currentGame)))
	minScore := math.MaxInt32

	for _, cellChange := range game.GetAvailableCellChanges(currentGame) {

		virtualGame, _ := game.PlayTurn(currentGame, cellChange)
		reversePlayerScore := GetMaxScore(virtualGame, aiPlayer, depth+1, depthLimit)

		if reversePlayerScore < minScore {
			minScore = reversePlayerScore
		}

	}

	return minScore

}

func GetBestCellChange(currentGame game.Game, aiPlayer player.Player, depth int, depthLimit int) cell.Cell {

	maxScore := -math.MaxInt32
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

	// maxScore := -math.MaxInt32
	// BestCellChange := cell.Cell{}
	//
	// maxScoreJobs := make(chan maxScoreJob, 100)
  // scores := make(chan scoreForCellChange, 100)
	// maxScoreWorkerCount := 2
	//
	// for workerIndex := 0; workerIndex <= maxScoreWorkerCount; workerIndex++ {
	// 	go GetMaxScoreWorker(maxScoreJobs, scores)
	// }
	//
	// availableCellChanges := game.GetAvailableCellChanges(currentGame)
	//
	// for _, cellChange := range availableCellChanges {
	//
	// 	virtualGame, _ := game.PlayTurn(currentGame, cellChange)
	// 	maxScoreJobs <- maxScoreJob{virtualGame, aiPlayer, depth, depthLimit}
	//
	// }
	//
	// close(maxScoreJobs)
	//
	// for cellChangeIdx := 0; cellChangeIdx < len(availableCellChanges)-1; cellChangeIdx++ {
	// 		if curMaxScore := <-scores; curMaxScore > maxScore{
	// 			maxScore := curMaxScore
	// 			BestCellChange := scoreForCellChange.cellChange
	// 		}
  // }
	//
	// return cell.Cell{}

}

// func GetMaxScoreWorker(maxScoreJobs <-chan maxScoreJob, scores chan<- scoreForCellChange){
//
// 		for job := range maxScoreJobs {
// 			scores <- GetMaxScore(job.Game, job.Player, job.Depth, job.DepthLimit)
// 		}
//
// }
//
// type maxScoreJob struct{
// 	Game game.Game
// 	Player player.Player
// 	Depth int
// 	DepthLimit int
// }
//
// type scoreForCellChange struct{
// 	score int
// 	cellChange cell.Cell
// }

func Score(gameBoard board.Board, aiPlayer player.Player, depth int) int {

	cellDist := board.GetCellDistribution(gameBoard)
	reverseCellType := cell.GetReverseCellType(aiPlayer.CellType)

	var winScore int

	if cellDist[aiPlayer.CellType] > cellDist[reverseCellType] {
		winScore = math.MaxInt32 - depth
	} else if cellDist[aiPlayer.CellType] < cellDist[reverseCellType] {
		winScore = -math.MaxInt32 + depth
	} else {
		winScore = 0
	}

	//fmt.Println("__debug__")
	// Enhance with "techniques particulières à Othello"
	// http://www.ffothello.org/informatique/algorithmes/

	return winScore

}
