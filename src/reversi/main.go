package main

import (
  "fmt"
  "reversi/game/game"
  "reversi/game/player"
  "reversi/game/cell"
)

func main(){

  fmt.Println("\n########## GAME STARTED ##########")

  playerBlack := askForPlayer("\n### Black player ###\n", cell.TypeBlack)
  playerWhite := askForPlayer("\n### White player ###\n", cell.TypeWhite)

  party := game.New(8, 8, []player.Player{playerBlack, playerWhite})

  fmt.Println(game.Render(party))

  for !game.IsFinished(party){

    currPlayer := game.GetCurrentPlayer(party)
    fmt.Printf("%s, It's our turn !", currPlayer.Name)



  }

}

func askForPlayer(header string, cellType uint8) player.Player{

  var isHuman string
  var name string

  fmt.Println(header)
  fmt.Print("Are you an human ? (y/n): ")
  fmt.Scanf("%s", &isHuman)

  if isHuman == "y" || isHuman == ""{
    fmt.Print("What's your name ?: ")
    fmt.Scanf("%s", &name)
    return player.New(name, player.BrainTypeHuman, cellType)
  }

  return player.New("Computer", player.BrainTypeAI, cellType)

}
