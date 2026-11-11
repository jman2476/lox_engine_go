package main

import (
	"fmt"
	"strings"

	imath "github.com/jman2476/lox_engine_go/internal/intmath"
)

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
				board[val] = 'N'
			} else {
				board[val] = 'n'
			}
		}
	}

	// slices.Reverse(board)

	var printBoard string
	var line strings.Builder
	var files = "  a b c d e f g h"

	for i := range 64 {
		if i%8 == 0 {
			line.Reset()
			rank := string(rune(49 + i/8))
			lineStart := fmt.Sprintf("%s ", rank)
			line.WriteString(lineStart)
		}
		if board[i] == 0 {
			r, f := (i/8)%2, i%2
			if r == 0 {
				if f == 0 {
					line.WriteRune('=')
				} else {
					line.WriteRune('.')
				}
			} else {
				if f == 1 {
					line.WriteRune('=')
				} else {
					line.WriteRune('.')
				}
			}
		} else {
			line.WriteRune(board[i])
		}

		line.WriteRune(' ')

		if i%8 == 7 {
			printBoard = prependLine(
				printBoard, line.String())
		}
	}

	fmt.Println("Board state:")
	fmt.Println(prependLine(printBoard, files))
}

func prependLine(current, newLine string) string {
	return newLine + "\n" + current
}
