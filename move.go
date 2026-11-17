package main

import (
	"errors"

	imath "github.com/jman2476/lox_engine_go/internal/intmath"
)

var (
	ErrNoPieceFound = errors.New("no piece found to move")
)

type Move struct {
	from      int
	to        int
	promotion Piece
}

func findPiece(square int, pieces *PieceSet) (Piece, error) {
	v := imath.Exp2(uint64(square))
	types := []Piece{Pawn, King, Queen, Bishop, Knight, Rook}

	for _, t := range types {
		set := pieces.GetPiece(t)
		if *set&v != 0 {
			return t, nil
		}
	}

	return None, ErrNoPieceFound
}

func movePiece(move string, pieceSet *PieceSet) error {
	mv, err := parseMove(move)
	if err != nil {
		return err
	}

	piece, err := findPiece(mv.from, pieceSet)
	if err != nil {
		return err
	}

	mask := pieceSet.GetPiece(piece)

	if mv.promotion == 0 {
		*mask += imath.Exp2(uint64(mv.to)) - imath.Exp2(uint64(mv.from))

	}

	return nil
}

func unMovePiece(move string, pieceSet *PieceSet) error {
	mv, err := parseMove(move)
	if err != nil {
		return err
	}

	piece, err := findPiece(mv.to, pieceSet)
	if err != nil {
		return err
	}

	mask := pieceSet.GetPiece(piece)

	if mv.promotion == 0 {
		*mask += imath.Exp2(uint64(mv.from)) - imath.Exp2(uint64(mv.to))
	}
	return nil
}

func validateMove(move string, piece Piece, board Board) bool {
	switch piece {
	case Pawn:
		return validatePawn(move, board)
	}

	return false
}

func validatePawn(move string, board Board) bool {
	return true
}
