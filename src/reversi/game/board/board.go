package board

import (
	"errors"
	"reversi/game/cell"
	"reversi/game/matrix"
	"reversi/game/vector"
	"strconv"
)

type Board [][]uint8

func New(xSize uint8, ySize uint8) Board {
	board := Board{}
	for y := uint8(0); y < ySize; y++ {
		board = append(board, make([]uint8, xSize, xSize))
	}
	return board
}

func IsValidBoardSize(xSize int, ySize int) bool {
	return xSize%2 == 0 && ySize%2 == 0
}

func InitCells(board Board) (Board, error) {
	xSize, ySize := matrix.GetSize(board)
	if !IsValidBoardSize(xSize, ySize) {
		return board, errors.New("Invalid board Size, x/y dim must be even to place departure cells")
	}
	return DrawCells(GetDepartureCells(board), board), nil
}

func GetDepartureCells(board Board) []cell.Cell {

	xSize, ySize := matrix.GetSize(board)
	xMiddle := uint8(xSize / 2)
	yMiddle := uint8(ySize / 2)

	return []cell.Cell{
		cell.New(xMiddle, yMiddle, cell.TypeBlack),
		cell.New(xMiddle-1, yMiddle-1, cell.TypeBlack),
		cell.New(xMiddle-1, yMiddle, cell.TypeWhite),
		cell.New(xMiddle, yMiddle-1, cell.TypeWhite),
	}

}

func Render(board Board, cellProposals []cell.Cell) string {

	renderMatrix := [][]string{}

	for yPos, row := range board {
		renderMatrix = append(renderMatrix, make([]string, len(row)))
		for xPos, cellType := range row {
			_, proposalCellIdx := FindCellIndexAt(uint8(xPos), uint8(yPos), cellProposals)
			if proposalCellIdx != -1 {
				renderMatrix[yPos][xPos] = strconv.Itoa(proposalCellIdx)
			} else {
				renderMatrix[yPos][xPos] = cell.GetSymbol(cellType)
			}
		}
	}

	return matrix.Render(renderMatrix)

}

func IsFull(board Board) bool {
	for _, ySlice := range board {
		for _, cellType := range ySlice {
			if cellType == cell.TypeEmpty {
				return false
			}
		}
	}
	return true
}

func DrawCells(cells []cell.Cell, board Board) Board {
	newBoard := Copy(board)
	for _, cell := range cells {
		newBoard[cell.Y][cell.X] = cell.CellType
	}
	return newBoard
}

func Copy(srcBoard Board) Board {
	dstBoard := make(Board, len(srcBoard))
	for idx, row := range srcBoard {
		dstBoard[idx] = make([]uint8, len(row))
		copy(dstBoard[idx], srcBoard[idx])
	}
	return dstBoard
}

func GetCellType(xPos uint8, yPos uint8, board Board) uint8 {
	if !(uint8(len(board)-1) >= yPos && uint8(len(board[yPos])-1) >= xPos) {
		return cell.TypeEmpty
	}
	return board[yPos][xPos]
}

func GetFlippedCellsFromCellChange(cellChange cell.Cell, board Board) []cell.Cell {

	if GetCellType(cellChange.X, cellChange.Y, board) != cell.TypeEmpty {
		return []cell.Cell{}
	}

	flippedCells := []cell.Cell{}

	for _, directionalVector := range vector.GetDirectionalVectors() {
		flippedInDirection := GetFlippedCellsForCellChangeAndDirectionVector(cellChange, directionalVector, board)
		flippedCells = append(flippedCells, flippedInDirection...)
	}

	return flippedCells

}

func GetFlippedCellsForCellChangeAndDirectionVector(cellChange cell.Cell, directionVector vector.Vector, board Board) []cell.Cell {

	flippedCells := []cell.Cell{}

	var localCellType uint8
	localCellPosition := vector.Vector{int(cellChange.X), int(cellChange.Y)}
	reverseCellType := cell.GetReverseCellType(cellChange.CellType)

	for {
		localCellPosition = vector.VectorAdd(localCellPosition, directionVector)
		localCellType = GetCellType(uint8(localCellPosition.X), uint8(localCellPosition.Y), board)
		if localCellType != reverseCellType {
			break
		}
		flippedCell := cell.New(uint8(localCellPosition.X), uint8(localCellPosition.Y), cellChange.CellType)
		flippedCells = append(flippedCells, flippedCell)
	}

	if localCellType == cellChange.CellType && len(flippedCells) > 0 {
		return flippedCells
	}

	return []cell.Cell{}

}

func IsLegalCellChange(cellChange cell.Cell, board Board) bool {

	if GetCellType(cellChange.X, cellChange.Y, board) != cell.TypeEmpty {
		return false
	}

	for _, directionalVector := range vector.GetDirectionalVectors() {
		flippedInDirection := GetFlippedCellsForCellChangeAndDirectionVector(cellChange, directionalVector, board)
		if len(flippedInDirection) > 0 {
			return true
		}
	}

	return false

}

func GetLegalCellChangesForCellType(cellType uint8, board Board) []cell.Cell {

	legalCellChanges := []cell.Cell{}

	for y, row := range board {
		for x, _ := range row {
			cellChange := cell.Cell{uint8(x), uint8(y), cellType}
			if IsLegalCellChange(cellChange, board) {
				legalCellChanges = append(legalCellChanges, cellChange)
			}
		}
	}

	return legalCellChanges

}

func GetCellDistribution(board Board) map[uint8]uint8 {
	dist := map[uint8]uint8{cell.TypeEmpty: uint8(0), cell.TypeBlack: uint8(0), cell.TypeWhite: uint8(0)}
	for _, row := range board {
		for _, cellType := range row {
			dist[cellType]++
		}
	}
	return dist
}

func FindCellIndexAt(x uint8, y uint8, cells []cell.Cell) (cell.Cell, int) {
	for idx, cell := range cells {
		if cell.X == x && cell.Y == y {
			return cell, idx
		}
	}
	return cell.Cell{}, -1
}
