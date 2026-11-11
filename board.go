package main

import (
	imath "github.com/jman2476/lox_engine_go/internal/intmath"
)

type Board struct {
	White PieceSet
	Black PieceSet
}

type PieceSet struct {
	Pawn   uint64
	King   uint64
	Queen  uint64
	Bishop uint64
	Knight uint64
	Rook   uint64
}

type Side int

const (
	White Side = iota
	Black
)

func (b *Board) SetNew() {
	b.White = SetupSide(White)
	b.Black = SetupSide(Black)
}

func SetupSide(side Side) (pieces PieceSet) {
	pieces.Pawn = SetPawns(side)
	pieces.King = SetKings(side)
	pieces.Queen = SetQueens(side)
	pieces.Bishop = SetBishops(side)
	pieces.Knight = SetKnights(side)
	pieces.Rook = SetRooks(side)

	return
}

func SetPawns(side Side) (positions uint64) {
	// row := 0b11111111
	switch side {
	case White:
		for i := 8; i < 16; i++ {
			positions += imath.Exp2(uint64(i))
		}
	case Black:
		for i := 48; i < 56; i++ {
			positions += imath.Exp2(uint64(i))
		}
	}

	return
}

func SetKings(side Side) uint64 {
	switch side {
	case White:
		return imath.Exp2(4)
	case Black:
		return imath.Exp2(60)
	}

	// Error value
	return imath.Exp2(4) + imath.Exp2(60)
}

func SetQueens(side Side) uint64 {
	switch side {
	case White:
		return imath.Exp2(3)
	case Black:
		return imath.Exp2(59)
	}

	// Error value
	return imath.Exp2(3) + imath.Exp2(59)
}

func SetBishops(side Side) uint64 {
	switch side {
	case White:
		return imath.Exp2(2) + imath.Exp2(5)
	case Black:
		return imath.Exp2(58) + imath.Exp2(61)
	}

	return imath.Exp2(2) + imath.Exp2(5) + imath.Exp2(58) + imath.Exp2(61)
}

func SetKnights(side Side) uint64 {
	switch side {
	case White:
		return imath.Exp2(1) + imath.Exp2(6)
	case Black:
		return imath.Exp2(57) + imath.Exp2(62)
	}

	return imath.Exp2(2) + imath.Exp2(5) + imath.Exp2(58) + imath.Exp2(61)
}

func SetRooks(side Side) uint64 {
	switch side {
	case White:
		return imath.Exp2(0) + imath.Exp2(7)
	case Black:
		return imath.Exp2(56) + imath.Exp2(63)
	}

	return imath.Exp2(2) + imath.Exp2(5) + imath.Exp2(58) + imath.Exp2(61)
}

func (b *Board) GetSide(side Side) PieceSet {
	if side == White {
		return b.White
	}
	return b.Black
}
