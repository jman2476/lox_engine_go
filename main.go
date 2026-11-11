package main

func main() {
	var board Board
	board.White.Pawn = SetPawns(White)
	board.Black.Pawn = SetPawns(Black)

	board.Print()
}
