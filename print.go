package main

import (
	"fmt"
	"slices"
	"strings"

	imath "github.com/jman2476/lox_engine_go/internal/intmath"
)

type PrintBoard struct {
	a []rune
	b []rune
	c []rune
	d []rune
	e []rune
	f []rune
	g []rune
	h []rune
}

var (
	whitePieces = []rune{'K', 'Q', 'P', 'N', 'B', 'R'}
	blackPieces = []rune{'k', 'q', 'p', 'n', 'b', 'r'}
)

func (b *Board) Print() {
	board := make([]rune, 64)

	for _, side := range []Side{White, Black} {
		set := b.GetSide(side)
		pawnidx := imath.GetIndices(set.Pawn)
		for _, val := range pawnidx {
			if side == White {
				board[val] = 'P'
			} else {
				board[val] = 'p'
			}
		}

		kingidx := imath.GetIndices(set.King)
		for _, val := range kingidx {
			if side == White {
				board[val] = 'K'
			} else {
				board[val] = 'k'
			}
		}

		queenidx := imath.GetIndices(set.Queen)
		for _, val := range queenidx {
			if side == White {
				board[val] = 'Q'
			} else {
				board[val] = 'q'
			}
		}

		rookidx := imath.GetIndices(set.Rook)
		for _, val := range rookidx {
			if side == White {
				board[val] = 'R'
			} else {
				board[val] = 'r'
			}
		}

		bishopidx := imath.GetIndices(set.Bishop)
		for _, val := range bishopidx {
			if side == White {
				board[val] = 'B'
			} else {
				board[val] = 'b'
			}
		}

		knightidx := imath.GetIndices(set.Knight)
		for _, val := range knightidx {
			if side == White {
				board[val] = 'K'
			} else {
				board[val] = 'k'
			}
		}
	}

	slices.Reverse(board)

	var printBoard strings.Builder
	for i := range 64 {
		if i%8 == 0 {
			printBoard.WriteRune('\n')
		}
		printBoard.WriteRune(board[i])
	}

	fmt.Println("Board state:")
	fmt.Print(printBoard.String())
}
