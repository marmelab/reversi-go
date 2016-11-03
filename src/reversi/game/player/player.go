package player

type Player struct {
	Name        string
	HumanPlayer bool
	CellType    uint8
}

func New(name string, humanPlayer bool, cellType uint8) Player {
	return Player{
		Name:        name,
		HumanPlayer: humanPlayer,
		CellType:    cellType,
	}
}
