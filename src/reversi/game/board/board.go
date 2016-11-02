package board

import (
  "bytes"
  "strings"
  "reversi/game/cell"
  "github.com/fatih/color"
)

type Board [][]uint8

func New(xSize uint8, ySize uint8) Board{
  board := Board{}
  var y uint8
  for y = 0; y < ySize; y++{
    board = append(board, make([]uint8, xSize, xSize))
  }
  board = DrawCells(GetDepartureCells(board), board)
  return board
}

func GetDepartureCells(board Board) []cell.Cell{

  xSize, ySize := GetSize(board)

  xMiddle := uint8(xSize/2)
  yMiddle := uint8(ySize/2)

  return []cell.Cell{
    cell.New(xMiddle, yMiddle, cell.TypeBlack),
    cell.New(xMiddle - 1, yMiddle - 1, cell.TypeBlack),
    cell.New(xMiddle - 1, yMiddle, cell.TypeWhite),
    cell.New(xMiddle, yMiddle - 1, cell.TypeWhite),
  }

}

func Render(board Board) string{
  var buffer bytes.Buffer
  xSize, _ := GetSize(board)
  underlined := color.New(color.Underline).SprintFunc()
  buffer.WriteString(strings.Repeat("_", int((xSize * 2) + 1)) + "\n")
  for _, row := range board{
    buffer.WriteString("|")
    for _, cellType := range row{
      buffer.WriteString(underlined(cell.GetSymbol(cellType) + "|"))
    }
    buffer.WriteString("\n")
  }

  return buffer.String()
}

func IsFull(board Board) bool{
  for _, ySlice := range board{
    for _, cellType := range ySlice{
      if cellType == cell.TypeEmpty{
        return false
      }
    }
  }
  return true
}

func GetSize(board Board) (uint8, uint8){
  if len(board) == 0 {
    return 0, 0
  }
  return uint8(len(board[0])), uint8(len(board))
}

func DrawCells(cells []cell.Cell, board Board) Board{

  newBoard := board

  for _, cell := range cells{
    if CellExist(cell.X, cell.Y, newBoard){
      newBoard[cell.Y][cell.X] = cell.CellType
    }
  }

  return newBoard

}

func CellExist(xPos uint8, yPos uint8, board Board) bool{
  return uint8(len(board) - 1) > yPos && uint8(len(board[yPos]) - 1) > xPos
}

// func GetFlippedCellsFromCellChange(cell cell.Cell, board Board) []cell.Cell{
//
//   flipped = []cell.Cell
//
//   if !CellExist(cell.X, cell.Y, board) || board[cell.Y][cell.X] != cell.TypeEmpty{
//     return flipped
//   }
//
//
//
// }

func GetCellDistribution(board Board) map[uint8]uint8{
  dist := map[uint8]uint8{cell.TypeEmpty: uint8(0), cell.TypeBlack: uint8(0), cell.TypeWhite: uint8(0)}
  for _, row := range board{
    for _, cellType := range row{
      dist[cellType]++
    }
  }
  return dist
}
