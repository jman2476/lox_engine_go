package main

import (
	"errors"

	imath "github.com/jman2476/lox_engine_go/internal/intmath"
)

var (
	ErrNoPieceFound = errors.New("no piece found to move")
)

func findPiece(square int, pieces PieceSet) (rune, error) {
	v := imath.Exp2(uint64(square))
	types := []rune{'p', 'k', 'q', 'b', 'r', 'n'}

	for _, t := range types {
		set := pieces.GetPiece(t)
		if set&v != 0 {
			return t, nil
		}
	}

	return 0, ErrNoPieceFound
}

func movePiece(move string, pieces *PieceSet) error {

	return nil
}
