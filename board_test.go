package main

import "testing"

func TestNoOverlaps(t *testing.T) {
	// var sum uint64
	var board Board

	sum := board.bBishop & board.wBishop & board.wPawn & board.bPawn & board.wKing & board.bKing & board.wQueen & board.bQueen & board.wKnight & board.bKnight & board.wRook & board.bRook

	if sum != 0 {
		t.Errorf("Overlaps exist on the board: resulting value is %b", sum)
	}
}
