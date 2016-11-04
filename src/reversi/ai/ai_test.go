package ai

import (
	"fmt"
	"time"
	//"reversi/game/board"
	"reversi/game/cell"
	"reversi/game/game"
	"reversi/game/player"
	"testing"
)

func TestGetBestCellChange(t *testing.T) {

	playerBlack := player.New("John", true, cell.TypeBlack)
	playerWhite := player.New("Doe", true, cell.TypeWhite)

	party := game.New([]player.Player{playerBlack, playerWhite})
	//party.Board = board.Board{{0, 0, 1}, {1, 2, 2}, {0, 2, 2}, {0, 0, 0}}

	fmt.Println(game.RenderAskBoard(party))
	fmt.Println("----")

	start := time.Now()
	fmt.Println(GetBestCellChange(party, playerBlack, 0, 6))
	elapsed := time.Since(start)
	fmt.Printf("Calculation took %s", elapsed)
	fmt.Println("----")

}

func TestScore(t *testing.T) {

	// testBoard := board.Board{{0, 0, 0}, {2, 2, 0}, {0, 2, 1}}
	//
	// if Score(testBoard, player.New("Doe", true, cell.TypeBlack), 1) != -999 {
	// 	t.Error("Incorrect score for board and player")
	// }
	//
	// if Score(testBoard, player.New("Doe", true, cell.TypeWhite), 1) != 999 {
	// 	t.Error("Incorrect score for board and player")
	// }

}
