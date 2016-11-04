package game

import (
	"errors"
	"fmt"
	"reversi/game/board"
	"reversi/game/cell"
	"reversi/game/player"
	"strings"
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

func GetCurrentPlayer(game Game) player.Player {
	return game.Players[game.CurrPlayerIndex]
}

func GetScore(game Game) map[player.Player]uint8 {
	dist := board.GetCellDistribution(game.Board)
	score := make(map[player.Player]uint8, 2)
	for _, player := range game.Players {
		score[player] = dist[player.CellType]
	}
	return score
}

func SwitchPlayer(game Game) Game {

	newGame := game
	if newGame.CurrPlayerIndex == 0 {
		newGame.CurrPlayerIndex = 1
	} else {
		newGame.CurrPlayerIndex = 0
	}
	return newGame
}

func CanPlayerChangeCells(player player.Player, currentGame Game) bool {
	return len(board.GetLegalCellChangesForCellType(player.CellType, currentGame.Board)) > 0
}

func RenderAskBoard(game Game) string {
	currentPlayer := GetCurrentPlayer(game)
	legalCellChanges := board.GetLegalCellChangesForCellType(currentPlayer.CellType, game.Board)
	return board.Render(game.Board, legalCellChanges)
}

func PlayTurn(currentGame Game) (Game, error) {

	if !CanPlayerChangeCells(GetCurrentPlayer(currentGame), currentGame) {
		return SwitchPlayer(currentGame), errors.New("You can't play !")
	}

	newGame := currentGame
	cellChange := askForCellChange(newGame)

	cellChangesFromChoice := append(board.GetFlippedCellsFromCellChange(cellChange, newGame.Board), cellChange)
	newGame.Board = board.DrawCells(cellChangesFromChoice, newGame.Board)

	return SwitchPlayer(newGame), nil

}

func askForCellChange(game Game) cell.Cell {

	var legalCellChangeChoice int
	currentPlayer := GetCurrentPlayer(game)
	legalCellChanges := board.GetLegalCellChangesForCellType(currentPlayer.CellType, game.Board)

	fmt.Printf("%s, It's our turn !\n", strings.ToUpper(currentPlayer.Name))

	if currentPlayer.HumanPlayer {
		fmt.Printf("Which position do you choose (0..%d) ? ", len(legalCellChanges)-1)
		fmt.Scanf("%d\n", &legalCellChangeChoice)
	} else {
		legalCellChangeChoice = 0 // todo => AI
		fmt.Printf("AI makes his choice ! %d\n", legalCellChangeChoice)
	}

	return legalCellChanges[legalCellChangeChoice]

}
