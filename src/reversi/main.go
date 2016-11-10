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

const CENTERING_BOARD_SPACE_COUNT int = 12
const AI_REFLECTION_TIME time.Duration = time.Millisecond * 1500
const GAME_WIDTH = 45

func main() {

	fmt.Printf(Centering(" REVERSI ", GAME_WIDTH, "#"))

	playerBlack := AskForPlayer(Centering(" Black player ", GAME_WIDTH, "-"), cell.TypeBlack)
	playerWhite := AskForPlayer(Centering(" White player ", GAME_WIDTH, "-"), cell.TypeWhite)

	currentGame := game.New([]player.Player{playerBlack, playerWhite})

	fmt.Println(Centering(" GAME STARTED ", GAME_WIDTH, "#"))
	fmt.Println(IndentString(game.Render(currentGame), CENTERING_BOARD_SPACE_COUNT))

	var cellChange cell.Cell
	var playErr error

	for !game.IsFinished(currentGame) {

		currentPlayer := game.GetCurrentPlayer(currentGame)
		currentPlayerName := strings.ToUpper(currentPlayer.Name)
		currentPlayerSymbol := cell.GetSymbol(currentPlayer.CellType)

		if currentPlayer.HumanPlayer {
			fmt.Println(Centering(fmt.Sprintf(" %s (%s), It's our turn ! ", currentPlayerName, currentPlayerSymbol), GAME_WIDTH, "#"))
			fmt.Println(IndentString(game.RenderAskBoard(currentGame), CENTERING_BOARD_SPACE_COUNT))
			cellChange = game.AskForCellChange(currentGame)
		} else {
			fmt.Println(Centering(fmt.Sprintf(" It's the turn of %s (%s), ... ", currentPlayerName, currentPlayerSymbol), GAME_WIDTH, "#"))
			fmt.Println(IndentString(game.RenderAskBoard(currentGame), CENTERING_BOARD_SPACE_COUNT))
			cellChange, _ = ai.GetBestCellChangeInTime(currentGame.Board, currentPlayer.CellType, AI_REFLECTION_TIME)
			fmt.Println(Centering(fmt.Sprintf(" %s (%s) changes cell at %d, %d ", currentPlayerName, currentPlayerSymbol, cellChange.X+1, cellChange.Y+1), GAME_WIDTH, " "))
		}

		currentGame, playErr = game.PlayTurn(currentGame, cellChange)

		if !game.IsFinished(currentGame) && playErr != nil {
			fmt.Println(playErr)
			if _, ok := playErr.(game.NoPossibilityError); ok {
				break
			}
		}

	}

	fmt.Println(Centering(" END OF GAME ", GAME_WIDTH, "#"))
	fmt.Println(IndentString(game.Render(currentGame), CENTERING_BOARD_SPACE_COUNT))

	if winner, err := game.GetWinner(currentGame); err == nil {
		winnerNotice := fmt.Sprintf("\n########## %s (%s) WINS ! ##########\n\n", strings.ToUpper(winner.Name), cell.GetSymbol(winner.CellType))
		fmt.Println(Centering(winnerNotice, GAME_WIDTH, "#"))
	} else {
		fmt.Println(err)
	}

}

func IndentString(str string, spaceCount int) string {
	return strings.Repeat(" ", spaceCount) + strings.Replace(str, "\n", "\n"+strings.Repeat(" ", spaceCount), -1)
}

func Centering(message string, width int, spaceChar string) string {

	var headerDecoration string

	if len(message) > width {
		headerDecoration = ""
	} else {
		headerDecoration = strings.Repeat(spaceChar, (width-len(message))/2)
	}

	return strings.Join([]string{"\n", headerDecoration, message, headerDecoration, "\n"}, "")

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
