package matrix

import (
	"testing"
)

func TestGetSizeShouldReturnMatrixSizeFromFirstRow(t *testing.T) {

	matrix := [][]string{{"str1"}, {"str1"}}
	x, y := GetSize(matrix)

	if x != 1 || y != 2 {
		t.Error("The expected matrix size is not right")
	}
}
