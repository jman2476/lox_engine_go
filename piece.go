package main

type Piece uint8

const (
	None Piece = iota
	Pawn
	King
	Queen
	Bishop
	Knight
	Rook
)
