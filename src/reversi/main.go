package main

import (
	"fmt"
	"reversi/ai"
	"reversi/game/cell"
	"reversi/game/game"
	"reversi/game/player"
	"strings"
	"time"
)

const CENTERING_BOARD_SPACE_COUNT int = 7
const AI_REFLECTION_TIME time.Duration = time.Millisecond * 1500

func main() {

	fmt.Println("\n############# REVERSI #############")

	playerBlack := AskForPlayer("\n### Black player ###\n", cell.TypeBlack)
	playerWhite := AskForPlayer("\n### White player ###\n", cell.TypeWhite)

	currentGame := game.New([]player.Player{playerBlack, playerWhite})

	fmt.Println("\n########## INITIAL BOARD ##########")
	fmt.Println(IndentString(game.Render(currentGame), CENTERING_BOARD_SPACE_COUNT))
	fmt.Println("########## GAME STARTED ##########\n")

	var cellChange cell.Cell
	var playErr error

	for !game.IsFinished(currentGame) {

		currentPlayer := game.GetCurrentPlayer(currentGame)

		if currentPlayer.HumanPlayer {
			fmt.Printf("## %s (%s), It's our turn ! Make a choice ##\n", strings.ToUpper(currentPlayer.Name), cell.GetSymbol(currentPlayer.CellType))
			fmt.Println(IndentString(game.RenderAskBoard(currentGame), CENTERING_BOARD_SPACE_COUNT))
			cellChange = game.AskForCellChange(currentGame)
		} else {
			fmt.Printf("## It's the turn of %s (%s), ...  ##\n", strings.ToUpper(currentPlayer.Name), cell.GetSymbol(currentPlayer.CellType))
			fmt.Println(IndentString(game.RenderAskBoard(currentGame), CENTERING_BOARD_SPACE_COUNT))
			cellChange, _ = ai.GetBestCellChangeInTime(currentGame.Board, currentPlayer.CellType, AI_REFLECTION_TIME)
			fmt.Printf("%s (%s) changes cell at %d, %d\n\n", strings.ToUpper(currentPlayer.Name), cell.GetSymbol(currentPlayer.CellType), cellChange.X+1, cellChange.Y+1)
		}

		currentGame, playErr = game.PlayTurn(currentGame, cellChange)

		if !game.IsFinished(currentGame) && playErr != nil {
			fmt.Println(playErr)
			if _, ok := playErr.(game.NoPossibilityError); ok {
				break
			}
		}

	}

	fmt.Println("\n########## END OF GAME ##########\n")
	fmt.Println(IndentString(game.Render(currentGame), CENTERING_BOARD_SPACE_COUNT))

	if winner, err := game.GetWinner(currentGame); err == nil {
		fmt.Printf("\n########## %s (%s) WINS ! ##########\n\n", strings.ToUpper(winner.Name), cell.GetSymbol(winner.CellType))
	} else {
		fmt.Println(err)
	}

}

func IndentString(str string, spaceCount int) string {
	return strings.Repeat(" ", spaceCount) + strings.Replace(str, "\n", "\n"+strings.Repeat(" ", spaceCount), -1)
}

func AskForPlayer(header string, cellType uint8) player.Player {

	var isHumanInput string
	var name string

	fmt.Println(header)
	fmt.Print("Are you an human ? (y/n): ")
	fmt.Scanf("%s", &isHumanInput)
	isHuman := (isHumanInput == "y" || isHumanInput == "")

	if isHuman {
		fmt.Print("What's your name ?: ")
	} else {
		fmt.Print("What's the computer name ?: ")
	}

	fmt.Scanf("%s", &name)

	return player.New(name, isHuman, cellType)

}
