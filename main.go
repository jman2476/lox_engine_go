package main

func main() {
	VerifyStartingPositions()
	var board Board
	board.SetNew()

	board.Print()
}

func VerifyStartingPositions() {
	var board Board
	board.White.Pawn = SetPawns(White)
	board.Black.Pawn = SetPawns(Black)
	board.Print()
	board.White.Pawn = 0
	board.Black.Pawn = 0

	board.White.King = SetKings(White)
	board.Black.King = SetKings(Black)
	board.Print()
	board.White.King = 0
	board.Black.King = 0

	board.White.Queen = SetQueens(White)
	board.Black.Queen = SetQueens(Black)
	board.Print()
	board.White.Queen = 0
	board.Black.Queen = 0

	board.White.Bishop = SetBishops(White)
	board.Black.Bishop = SetBishops(Black)
	board.Print()
	board.White.Bishop = 0
	board.Black.Bishop = 0

	board.White.Rook = SetRooks(White)
	board.Black.Rook = SetRooks(Black)
	board.Print()
	board.White.Rook = 0
	board.Black.Rook = 0

	board.White.Knight = SetKnights(White)
	board.Black.Knight = SetKnights(Black)
	board.Print()
}
