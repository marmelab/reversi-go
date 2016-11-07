package ai

import (
	//"errors"
	"math"
	"reversi/game/board"
	"reversi/game/cell"
	//"reversi/game/game"
	//"reversi/game/matrix"
	//"reversi/game/player"
	"fmt"
	"time"
)

type Node struct {
	Board          board.Board
	RootCellChange cell.Cell
	IsOpponent     bool
	CellType       uint8
}

type ScoreNode struct {
	RootCellChange cell.Cell
	Score          int
}

func GetBestCellChangeInTime(currentBoard board.Board, cellType uint8) cell.Cell {

	nodes := make(chan Node)
	scores := make(chan ScoreNode)
	timeout := make(chan bool, 1)

	defer close(nodes)
	defer close(scores)

	go func() {
		time.Sleep(time.Second * 5)
		timeout <- true
	}()

	fmt.Println("start")

	go ProcessNodeScoring(Node{currentBoard, cell.Cell{}, false, cellType}, nodes)

	finished := false
	maxScore := -math.MaxInt32
	cellChange := cell.Cell{}

	for !finished {
		select {
		case finished = <-timeout:
		case node := <-nodes:
			fmt.Println("Receive Node")
			scores <- ScoreNode{node.RootCellChange, Score(node)}
		case score := <-scores:
			fmt.Println("Receive ScoreNode")
			if score.Score > maxScore {
				maxScore = score.Score
				cellChange = score.RootCellChange
			}
		}
	}

	return cellChange

}

func ProcessNodeScoring(rootNode Node, nodes chan Node) {

	fmt.Println("ProcessNode ", rootNode)

	var processNodes []Node

	// Scoring current node stage for each subnode
	for _, cellChange := range board.GetLegalCellChangesForCellType(rootNode.CellType, rootNode.Board) {
		node := Node{GetBoardFromCellChange(rootNode.Board, cellChange), cellChange, !rootNode.IsOpponent, cell.GetReverseCellType(rootNode.CellType)}
		nodes <- node
		processNodes = append(processNodes, node)
	}

	// Recursively scoring subnodes stages for each node at this stage
	for _, node := range processNodes {
		ProcessNodeScoring(node, nodes)
	}

}

func Score(node Node) int {
	return 0
}

func GetBoardFromCellChange(currentBoard board.Board, cellChange cell.Cell) board.Board {
	cellChangesToApply := append(board.GetFlippedCellsFromCellChange(cellChange, currentBoard), cellChange)
	return board.DrawCells(cellChangesToApply, currentBoard)
}

// const ScoringLevelLimit int = math.MaxInt8
//
// func GetMaxScore(currentGame game.Game, depth int, depthLimit int) int {
//
// 	if game.IsFinished(currentGame) || depth >= depthLimit {
// 		return Score(currentGame.Board, game.GetCurrentPlayer(currentGame), depth)
// 	}
//
// 	reversePlayerMaxScore := -ScoringLevelLimit
//
// 	for _, cellChange := range game.GetAvailableCellChanges(currentGame) {
//
// 		virtualGame, _ := game.PlayTurn(currentGame, cellChange)
// 		reversePlayerScore := GetMinScore(virtualGame, depth+1, depthLimit)
//
// 		if reversePlayerScore > reversePlayerMaxScore {
// 			reversePlayerMaxScore = reversePlayerScore
// 		}
//
// 	}
//
// 	return reversePlayerMaxScore
//
// }
//
// func GetMinScore(currentGame game.Game, depth int, depthLimit int) int {
//
// 	if game.IsFinished(currentGame) || depth >= depthLimit {
// 		return Score(currentGame.Board, game.GetCurrentPlayer(currentGame), depth)
// 	}
//
// 	reversePlayerMinScore := ScoringLevelLimit
//
// 	for _, cellChange := range game.GetAvailableCellChanges(currentGame) {
//
// 		virtualGame, _ := game.PlayTurn(currentGame, cellChange)
// 		reversePlayerScore := GetMaxScore(virtualGame, depth+1, depthLimit)
//
// 		if reversePlayerScore < reversePlayerMinScore {
// 			reversePlayerMinScore = reversePlayerScore
// 		}
//
// 	}
//
// 	return reversePlayerMinScore
//
// }
//
// func GetBestCellChange(currentGame game.Game, depth int, depthLimit int) (cell.Cell, error) {
//
// 	maxScore := -ScoringLevelLimit
// 	bestCellChange := cell.Cell{}
//
// 	availableCellChanges := game.GetAvailableCellChanges(currentGame)
//
// 	if len(availableCellChanges) == 0 {
// 		return bestCellChange, errors.New("AI can't play!")
// 	}
//
// 	if len(availableCellChanges) == 1 {
// 		return availableCellChanges[0], nil
// 	}
//
// 	for _, cellChange := range availableCellChanges {
//
// 		virtualGame, playTurnError := game.PlayTurn(currentGame, cellChange)
// 		cellChangeScore := GetMaxScore(virtualGame, depth, depthLimit)
//
// 		if playTurnError != nil {
// 			return bestCellChange, playTurnError
// 		}
//
// 		if cellChangeScore > maxScore {
// 			maxScore = cellChangeScore
// 			bestCellChange = cellChange
// 		}
//
// 	}
//
// 	return bestCellChange, nil
//
// }
//
// func Score(gameBoard board.Board, gamePlayer player.Player, depth int) int {
//
// 	// Enhance with "techniques particulières à Othello"
// 	// http://www.ffothello.org/informatique/algorithmes/
//
// 	availableCellChanges := board.GetLegalCellChangesForCellType(gamePlayer.CellType, gameBoard)
//
// 	supremacyScore := GetSupremacyScore(gameBoard, gamePlayer.CellType, depth)
// 	zoningScore := GetZoningScore(availableCellChanges, gameBoard)
//
// 	return supremacyScore + zoningScore
//
// }
//
// func GetZoningScore(availableCellChanges []cell.Cell, gameBoard board.Board) int {
//
// 	zoningScore := 0
// 	xSize, ySize := matrix.GetSize(gameBoard)
//
// 	// Scoring Strategy
// 	// +1000 for board limits
// 	// +1500 for board corners
//
// 	for _, cellChange := range availableCellChanges {
// 		xPos := int(cellChange.X)
// 		yPos := int(cellChange.Y)
// 		if xPos == 0 || xPos == xSize-1 || yPos == 0 || yPos == ySize-1 {
// 			zoningScore += 1000
// 			if (yPos == 0 && xPos == 0) || (yPos == ySize-1 && xPos == xSize-1) || (yPos == ySize-1 && xPos == 0) || (yPos == 0 && xPos == xSize-1) {
// 				zoningScore += 500
// 			}
// 		}
// 	}
//
// 	return zoningScore
//
// }
//
// func GetSupremacyScore(gameBoard board.Board, cellType uint8, depth int) int {
//
// 	cellDist := board.GetCellDistribution(gameBoard)
// 	reverseCellType := cell.GetReverseCellType(cellType)
//
// 	// Score based on the number of cell of the player's cellType
// 	// The depth parameter permit to highlight near distribution configurations
//
// 	if cellDist[cellType] > cellDist[reverseCellType] {
// 		return ScoringLevelLimit - depth
// 	}
//
// 	if cellDist[cellType] < cellDist[reverseCellType] {
// 		return -ScoringLevelLimit + depth
// 	}
//
// 	return 0
//
// }
