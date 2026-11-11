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

	printBoard.WriteString("  1 2 3 4 5 6 7 8")

	for i := range 64 {
		if i%8 == 0 {
			rank := string(rune(104 - i/8))
			lineStart := fmt.Sprintf("\n%s ", rank)
			printBoard.WriteString(lineStart)
		}
		if board[i] == 0 {
			r, f := (i/8)%2, i%2
			if r == 0 {
				if f == 0 {
					printBoard.WriteRune('O')
				} else {
					printBoard.WriteRune('X')
				}
			} else {
				if f == 1 {
					printBoard.WriteRune('O')
				} else {
					printBoard.WriteRune('X')
				}
			}
		} else {
			printBoard.WriteRune(board[i])
		}
		printBoard.WriteRune(' ')
	}

	fmt.Println("Board state:")
	fmt.Println(printBoard.String())
}
