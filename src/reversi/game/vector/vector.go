package vector

type Vector struct {
	X float64
	Y float64
}

func VectorAdd(vector Vector, addVector Vector) Vector {
	return Vector{vector.X + addVector.X, vector.Y + addVector.Y}
}

func GetDirectionnalVectors() []Vector {

	return []Vector{
		Vector{0, 1},
		Vector{1, 1},
		Vector{1, 0},
		Vector{1, -1},
		Vector{0, -1},
		Vector{-1, -1},
		Vector{-1, 0},
		Vector{-1, 1},
	}

}
