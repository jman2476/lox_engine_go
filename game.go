package main

type Game struct {
	Board   Board
	Ply     uint16
	Turn    bool
	History []PastMove
}

type PastMove struct {
	boardMask uint16
	move      string
}
