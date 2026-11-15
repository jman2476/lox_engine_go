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

func (p Piece) String() string {
	switch p {
	case Pawn:
		return "pawn"
	case King:
		return "king"
	case Queen:
		return "queen"
	case Bishop:
		return "bishop"
	case Knight:
		return "knight"
	case Rook:
		return "rook"
	default:
		return "none"
	}
}
