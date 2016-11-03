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
	return board.IsFull(game.Board) || board.NoBodyCanApplyCellChange(game.Board)
}

func GetCurrentPlayer(game Game) player.Player {
	return game.Players[game.CurrPlayerIndex]
}

func GetReversePlayer(game Game) player.Player {
	return game.Players[GetReversePlayerIndex(game)]
}

func GetScores(game Game) map[player.Player]uint8 {
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

func GetReversePlayerIndex(game Game) uint8 {
	return uint8(math.Abs(float64(game.CurrPlayerIndex) - 1))
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

	cellChange := askForCellChange(newGame)
	return SwitchPlayer(PlayCellChange(currentGame, cellChange)), nil
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

	var legalCellChangeChoice int
	currentPlayer := GetCurrentPlayer(game)
	availableCellChanges := GetAvailableCellChanges(game)

	fmt.Printf("%s, It's our turn !\n", strings.ToUpper(currentPlayer.Name))

	if currentPlayer.HumanPlayer {
		fmt.Printf("Which position do you choose (0..%d) ? ", len(availableCellChanges)-1)
		fmt.Scanf("%d\n", &legalCellChangeChoice)
	} else {
		legalCellChangeChoice = 0 // todo => AI
		fmt.Printf("AI makes his choice ! %d\n", legalCellChangeChoice)
	}

	return availableCellChanges[legalCellChangeChoice]

}
