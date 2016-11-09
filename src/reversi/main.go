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

func main() {

	fmt.Println("\n############# REVERSI #############")

	playerBlack := askForPlayer("\n### Black player ###\n", cell.TypeBlack)
	playerWhite := askForPlayer("\n### White player ###\n", cell.TypeWhite)

	currentGame := game.New([]player.Player{playerBlack, playerWhite})

	fmt.Println("\n########## INITIAL BOARD ##########")
	fmt.Println(game.Render(currentGame))
	fmt.Println("\n########## GAME STARTED ##########")

	var cellChange cell.Cell
	var playErr error

	for !game.IsFinished(currentGame) {

		fmt.Println(game.RenderAskBoard(currentGame))

		currentPlayer := game.GetCurrentPlayer(currentGame)

		if currentPlayer.HumanPlayer {
			fmt.Printf("%s (%s), It's our turn !\n", strings.ToUpper(currentPlayer.Name), cell.GetSymbol(currentPlayer.CellType))
			cellChange = game.AskForCellChange(currentGame)
		} else {
			fmt.Printf("%s (%s) thinks about best positions..\n", strings.ToUpper(currentPlayer.Name), cell.GetSymbol(currentPlayer.CellType))
			cellChange, _ = ai.GetBestCellChangeInTime(currentGame.Board, currentPlayer.CellType, time.Millisecond*1500)
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
	fmt.Println(game.Render(currentGame))

	if winner, err := game.GetWinner(currentGame); err == nil {
		fmt.Printf("\n########## %s (%s) WINS ! ##########\n\n", strings.ToUpper(winner.Name), cell.GetSymbol(winner.CellType))
	} else {
		fmt.Println(err)
	}

}

func askForPlayer(header string, cellType uint8) player.Player {

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
