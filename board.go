package main

type Board struct {
	wPawn   uint64
	bPawn   uint64
	wKing   uint64
	bKing   uint64
	wQueen  uint64
	bQueen  uint64
	wRook   uint64
	bRook   uint64
	wBishop uint64
	bBishop uint64
	wKnight uint64
	bKnight uint64
}

const (
	wPawnStart uint64 = 2 * 2 * 2 * 2 * 2 * 2 * 2 * 2
)

func (b Board) Print() {

}
