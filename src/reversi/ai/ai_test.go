package ai

import (
	"fmt"
	"reversi/game/cell"
	"reversi/game/game"
	"reversi/game/player"
	"testing"
)

func TestGetBestCellChange(t *testing.T) {

	playerBlack := player.New("John", true, cell.TypeBlack)
	playerWhite := player.New("Doe", true, cell.TypeWhite)

	party := game.New([]player.Player{playerBlack, playerWhite})

	fmt.Println(game.RenderAskBoard(party))
	fmt.Println("----")

	cellChangeChannel := make(chan cell.Cell, 1)
	go BestCellChange(party, playerBlack, 0, 6, cellChangeChannel)


	fmt.Println("----")

}
