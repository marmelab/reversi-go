package main

import (
	"fmt"
	"reversi/game/cell"
	"reversi/game/game"
	"reversi/game/player"
)

func main() {

	fmt.Println("\n########## GAME STARTED ##########")

	playerBlack := askForPlayer("\n### Black player ###\n", cell.TypeBlack)
	playerWhite := askForPlayer("\n### White player ###\n", cell.TypeWhite)

	party := game.New([]player.Player{playerBlack, playerWhite})

	fmt.Println(game.Render(party))

	for !game.IsFinished(party) {

		currPlayer := game.GetCurrentPlayer(party)
		fmt.Printf("%s, It's our turn !", currPlayer.Name)

	}

}

func askForPlayer(header string, cellType uint8) player.Player {

	var isHuman string
	var name string

	fmt.Println(header)
	fmt.Print("Are you an human ? (y/n): ")
	fmt.Scanf("%s", &isHuman)

	if isHuman == "y" || isHuman == "" {
		fmt.Print("What's your name ?: ")
		fmt.Scanf("%s", &name)
		return player.New(name, true, cellType)
	}

	return player.New("Computer", false, cellType)

}
