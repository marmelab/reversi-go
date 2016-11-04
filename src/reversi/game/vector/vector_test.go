package vector

import (
	"testing"
)

func TestVectorAddShouldAddVectors(t *testing.T) {

	if VectorAdd(Vector{1, 0}, (Vector{3, 42})) != (Vector{4, 42}) {
		t.Error("VectorAdd Should add vectors")
	}

	if VectorAdd(Vector{1, 3}, (Vector{-1, -10})) != (Vector{0, -7}) {
		t.Error("VectorAdd Should work with negative vectors")
	}

}
