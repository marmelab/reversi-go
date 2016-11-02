package player

const BrainTypeHuman uint8 = 0
const BrainTypeAI uint8 = 1

type Player struct{
  Name      string
  BrainType uint8
  CellType  uint8
}

func New(name string, brainType uint8, cellType uint8) Player{
  return Player{
    Name: name,
    BrainType: brainType,
    CellType: cellType,
  }
}
