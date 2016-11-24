package node

import(
    "reversi/game/board"
    "reversi/game/cell"
)

type Node struct {
	IsOpponent	bool
	Parent		*Node
	Children	[]*Node
    Score       *int
	Board       board.Board
	CellChange  cell.Cell
    CellType    uint8
}

func New(isOpponent bool, currBoard board.Board, currCellChange cell.Cell, cellType uint8) Node{
    return Node{IsOpponent: isOpponent, Board: currBoard, CellChange: currCellChange, CellType: cellType}
}

func (node *Node) Add(currBoard board.Board, currCellChange cell.Cell) *Node {
	childNode := Node{IsOpponent: !node.IsOpponent, Parent: node, Board: currBoard, CellChange: currCellChange, CellType: cell.GetReverseCellType(node.CellType)}
	node.Children = append(node.Children, &childNode)
	return &childNode
}

func (node *Node) GetRootNode() *Node{
    if(node.Parent != nil){
        return node.Parent.GetRootNode()
    }
    return node
}

func (node *Node) Evaluate() {
	for _, cn := range node.Children {
		if len(node.Children) > 0 {
			cn.Evaluate()
		}

		if cn.Parent.Score == nil {
			cn.Parent.Score = cn.Score
		} else if cn.IsOpponent && cn.Score != nil && *cn.Score > *cn.Parent.Score {
			cn.Parent.Score = cn.Score
		} else if !cn.IsOpponent && cn.Score != nil && *cn.Score < *cn.Parent.Score {
			cn.Parent.Score = cn.Score
		}
	}
}
