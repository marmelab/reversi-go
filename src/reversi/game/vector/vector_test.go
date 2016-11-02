package vector

import (
	"testing"
)

func TestVectorAddShouldAddVectors(t *testing.T) {

	additionnalVector := Vector{3, 42}
	expectedVector := Vector{4, 42}

	if VectorAdd(Vector{1, 0}, additionnalVector) != expectedVector {
		t.Error("VectorAdd Should add vectors")
	}

}
