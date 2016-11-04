package cell

const TypeEmpty uint8 = 0
const TypeBlack uint8 = 1
const TypeWhite uint8 = 2

type Cell struct {
	X        uint8
	Y        uint8
	CellType uint8
}

func New(x uint8, y uint8, cellType uint8) Cell {
	return Cell{x, y, cellType}
}

func GetSymbol(cellType uint8) string {
	switch cellType {
	case TypeBlack:
		return "○"
	case TypeWhite:
		return "●"
	default:
		return " "
	}
}

func GetReverseCellType(cellType uint8) uint8 {
	if cellType == TypeBlack {
		return TypeWhite
	}
	if cellType == TypeEmpty {
		return TypeEmpty
	}
	return TypeBlack
}
