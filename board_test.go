package main

import "testing"

func TestNoOverlaps(t *testing.T) {
	// var sum uint64
	var board Board

	board.SetNew()

	white := board.White.Bishop & board.White.King & board.White.Knight & board.White.Queen & board.White.Rook & board.White.Pawn
	black := board.Black.Bishop & board.Black.King & board.Black.Knight & board.Black.Queen & board.Black.Rook & board.Black.Pawn

	if black&white != 0 {
		t.Errorf("Overlaps exist on the board: resulting value is %b", black&white)
	}
}
