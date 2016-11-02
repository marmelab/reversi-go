package game

import (
  "reversi/game/board"
  "reversi/game/player"
)

type Game struct{
  Board           board.Board
  Players         []player.Player
  CurrPlayerIndex uint8
}

func New(xSize uint8, ySize uint8, players []player.Player) Game{
  return Game{
    board.New(xSize, ySize),
    players,
    0,
  }
}

func Render(game Game) string{
  return board.Render(game.Board)
}

func IsFinished(game Game) bool{
  return board.IsFull(game.Board)
}

func GetCurrentPlayer(game Game) player.Player{
  return game.Players[game.CurrPlayerIndex]
}

func GetScore(game Game) map[player.Player]uint8{

  dist := board.GetCellDistribution(game.Board)
  score := make(map[player.Player]uint8, 2)

  for _, player := range game.Players{
    score[player] = dist[player.CellType]
  }

  return score

}
