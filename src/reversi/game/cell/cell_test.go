package cell

import (
	"testing"
)

func TestNewCellShouldReturnNewCell(t *testing.T) {

	cell := New(0, 0, TypeBlack)
	expectedCell := Cell{0, 0, TypeBlack}

	if cell != expectedCell {
		t.Error("New doesn't return expected cell struct")
	}

}

func TestGetSymbolShouldReturnAppropriateCellTypeSymbol(t *testing.T) {

	if GetSymbol(TypeBlack) != "○" {
		t.Error("Bad symbol for Black cellType")
	}

	if GetSymbol(TypeWhite) != "●" {
		t.Error("Bad symbol for White cellType")
	}

	if GetSymbol(TypeEmpty) != " " {
		t.Error("Bad symbol for empty cellType")
	}

}
