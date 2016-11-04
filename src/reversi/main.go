package main

import (
	"fmt"
	"reversi/ai"
	"reversi/game/cell"
	"reversi/game/game"
	"reversi/game/player"
	"strings"
)

func main() {

	fmt.Println("\n############# REVERSI #############")

	playerBlack := askForPlayer("\n### Black player ###\n", cell.TypeBlack)
	playerWhite := askForPlayer("\n### White player ###\n", cell.TypeWhite)

	party := game.New([]player.Player{playerBlack, playerWhite})

	fmt.Println("\n########## INITIAL BOARD ##########")
	fmt.Println(game.Render(party))
	fmt.Println("\n########## GAME STARTED ##########")

	var cellChange cell.Cell
	var err error

	for !game.IsFinished(party) {

		fmt.Println(game.RenderAskBoard(party))

		currentPlayer := game.GetCurrentPlayer(party)

		if currentPlayer.HumanPlayer {
			fmt.Printf("%s, It's our turn !\n", strings.ToUpper(currentPlayer.Name))
			cellChange = game.AskForCellChange(party)
		} else {
			fmt.Printf("%s thinks about best positions..\n", strings.ToUpper(currentPlayer.Name))
			cellChange = ai.GetBestCellChange(party, currentPlayer, 0, 5)
		}

		party, err = game.PlayTurn(party, cellChange)

		if err != nil {
			fmt.Println(err)
		}

	}

	fmt.Println("\n########## END OF GAME ##########\n")

	game.Render(party)

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
