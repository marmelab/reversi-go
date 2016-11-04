package game

import (
	"errors"
	"fmt"
	"reversi/game/board"
	"reversi/game/cell"
	"reversi/game/player"
)

type Game struct {
	Board           board.Board
	Players         []player.Player
	CurrPlayerIndex uint8
}

func New(players []player.Player) Game {
	gameBoard, _ := board.InitCells(board.New(8, 8))
	return Game{
		gameBoard,
		players,
		0,
	}
}

func Render(game Game) string {
	return board.Render(game.Board, []cell.Cell{})
}

func IsFinished(game Game) bool {
	return board.IsFull(game.Board)
}

func NoBodyCanApplyCellChange(game Game) bool {
	return !CanPlayerChangeCells(GetCurrentPlayer(game), game) && !CanPlayerChangeCells(GetReversePlayer(game), game)
}

func GetCurrentPlayer(game Game) player.Player {
	return game.Players[game.CurrPlayerIndex]
}

func GetReversePlayer(game Game) player.Player {
	return game.Players[GetReversePlayerIndex(game)]
}

func GetReversePlayerIndex(game Game) uint8 {
	if game.CurrPlayerIndex == 0 {
		return 1
	}
	return 0
}

func SwitchPlayer(game Game) Game {
	newGame := game
	newGame.CurrPlayerIndex = GetReversePlayerIndex(newGame)
	return newGame
}

func GetScores(game Game) map[player.Player]uint8 {
	dist := board.GetCellDistribution(game.Board)
	score := make(map[player.Player]uint8, 2)
	for _, player := range game.Players {
		score[player] = dist[player.CellType]
	}
	return score
}

func GetWinPlayer(game Game) (player.Player, error) {
	scores := GetScores(game)
	currentPlayerScore := scores[GetCurrentPlayer(game)]
	reversePlayerScore := scores[GetReversePlayer(game)]
	if currentPlayerScore > reversePlayerScore {
		return GetCurrentPlayer(game), nil
	}
	if reversePlayerScore > currentPlayerScore {
		return GetReversePlayer(game), nil
	}
	return player.Player{}, errors.New("There's no winner")
}

func CanPlayerChangeCells(player player.Player, currentGame Game) bool {

	return len(board.GetLegalCellChangesForCellType(player.CellType, currentGame.Board)) > 0
}

func RenderAskBoard(game Game) string {
	currentPlayer := GetCurrentPlayer(game)
	legalCellChanges := board.GetLegalCellChangesForCellType(currentPlayer.CellType, game.Board)
	return board.Render(game.Board, legalCellChanges)
}

func PlayTurn(currentGame Game, cellChange cell.Cell) (Game, error) {

	newGame := PlayCellChange(currentGame, cellChange)

	if !CanPlayerChangeCells(GetReversePlayer(newGame), newGame) {
		return newGame, errors.New("Opponent can't play ! Play Again !")
	} else if !CanPlayerChangeCells(GetCurrentPlayer(newGame), newGame) {
		return newGame, errors.New("There's no cell to play.")
	} else {
		return SwitchPlayer(newGame), nil
	}

}

func PlayCellChange(game Game, cellChange cell.Cell) Game {
	cellChanges := append(board.GetFlippedCellsFromCellChange(cellChange, game.Board), cellChange)
	return Game{
		board.DrawCells(cellChanges, game.Board),
		game.Players,
		game.CurrPlayerIndex,
	}
}

func GetAvailableCellChanges(game Game) []cell.Cell {
	return board.GetLegalCellChangesForCellType(GetCurrentPlayer(game).CellType, game.Board)
}

func AskForCellChange(game Game) cell.Cell {

	legalCellChangeChoice := 999
	availableCellChanges := GetAvailableCellChanges(game)

	for legalCellChangeChoice > len(availableCellChanges)-1 || legalCellChangeChoice < 0 {
		fmt.Printf("Which position do you choose (0..%d) ? \n", len(availableCellChanges)-1)
		fmt.Scanf("%d\n", &legalCellChangeChoice)
	}

	return availableCellChanges[legalCellChangeChoice]

}
