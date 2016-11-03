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
	var y uint8
	for y = 0; y < ySize; y++ {
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
			cellFinded, cellFindedIdx := cell.CellsContainsCellPosition(cell.New(uint8(xPos), uint8(yPos), cell.TypeEmpty), cellProposals)
			if cellFinded {
				renderMatrix[yPos][xPos] = strconv.Itoa(cellFindedIdx)
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
	newBoard := board
	for _, cell := range cells {
		newBoard[cell.Y][cell.X] = cell.CellType
	}
	return newBoard
}

func CellExist(xPos uint8, yPos uint8, board Board) bool {
	return uint8(len(board)-1) >= yPos && uint8(len(board[yPos])-1) >= xPos
}

func GetCellType(xPos uint8, yPos uint8, board Board) uint8 {
	if !CellExist(xPos, yPos, board) {
		return cell.TypeEmpty
	}
	return board[yPos][xPos]
}

func GetFlippedCellsFromCellChange(cellChange cell.Cell, board Board) []cell.Cell {

	cellChangeType := cellChange.CellType
	reverseCellType := cell.GetReverseCellType(cellChangeType)

	if GetCellType(cellChange.X, cellChange.Y, board) != cell.TypeEmpty {
		return []cell.Cell{}
	}

	var localCellType uint8
	var localVectorPosition vector.Vector

	flippedCells := []cell.Cell{}
	localFlippedCells := []cell.Cell{}

	for _, directionnalAddVector := range vector.GetDirectionnalVectors() {
		localFlippedCells = localFlippedCells[:0]
		localVectorPosition = vector.Vector{float64(cellChange.X), float64(cellChange.Y)}
		for {
			localVectorPosition = vector.VectorAdd(localVectorPosition, directionnalAddVector)
			localCellType = GetCellType(uint8(localVectorPosition.X), uint8(localVectorPosition.Y), board)
			if localCellType != reverseCellType {
				break
			}
			localCellChange := cell.New(uint8(localVectorPosition.X), uint8(localVectorPosition.Y), cellChangeType)
			localFlippedCells = append(localFlippedCells, localCellChange)
		}

		if localCellType == cellChangeType && len(localFlippedCells) > 0 {
			flippedCells = append(flippedCells, localFlippedCells...)
		}

	}

	return flippedCells

}

func IsLegalCellChange(cellChange cell.Cell, board Board) bool {
	return len(GetFlippedCellsFromCellChange(cellChange, board)) > 0
}

func GetLegalCellChangesForCellType(cellType uint8, board Board) []cell.Cell {

	legalCellChanges := []cell.Cell{}
	playableCells := GetPlayableCellsFromBoardByCellType(cellType, board)

	for _, playableCell := range playableCells {
		if IsLegalCellChange(playableCell, board) {
			legalCellChanges = append(legalCellChanges, playableCell)
		}
	}

	return legalCellChanges

}

func CellTypeHasCellChanges(cellType uint8, board Board) bool {
	return len(GetLegalCellChangesForCellType(cellType, board)) > 0
}

func GetPlayableCellsFromBoardByCellType(cellType uint8, board Board) []cell.Cell {

	// Todo => convolution matrix
	// stepSize := 2
	xSize, ySize := matrix.GetSize(board)
	// reverseCellType = cell.GetReverseCellType(cellType)
	stepSize := 1
	playableCells := []cell.Cell{}

	for yPos := 0; yPos < ySize; yPos += stepSize {
		for xPos := 0; xPos < xSize; xPos += stepSize {
			playableCells = append(playableCells, cell.New(uint8(xPos), uint8(yPos), cellType))
		}
	}

	return playableCells

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
