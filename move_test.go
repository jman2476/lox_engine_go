package main

import (
	"fmt"
	"testing"

	imath "github.com/jman2476/lox_engine_go/internal/intmath"
)

func TestFindPiece(t *testing.T) {
	var board Board
	board.SetNew()
	cases := []struct {
		input     int
		side      Side
		expected  rune
		expectErr error
	}{
		{
			input:     0,
			side:      White,
			expected:  'r',
			expectErr: nil,
		}, {
			input:     59,
			side:      Black,
			expected:  'q',
			expectErr: nil,
		}, {
			input:     32,
			side:      White,
			expected:  0,
			expectErr: ErrNoPieceFound,
		}, {
			input:     4,
			side:      Black,
			expected:  0,
			expectErr: ErrNoPieceFound,
		},
	}

	for _, c := range cases {
		var pieces PieceSet
		if c.side == White {
			pieces = board.White
		} else {
			pieces = board.Black
		}
		result, err := findPiece(c.input, &pieces)
		if err != nil && err != c.expectErr {
			t.Errorf(
				"Fail: Error\nCase: %d %v\nExpected: %v\nActual: %v",
				c.input, c.side,
				c.expectErr,
				err,
			)
		}

		if result != c.expected {
			t.Errorf(
				"Fail: Result\nCase: %d %v\nExpected: %d\nActual: %d",
				c.input, c.side,
				c.expected,
				result,
			)
		}
	}
}

func TestMovePiece(t *testing.T) {
	var board Board
	board.SetNew()

	cases := []struct {
		move     string
		piece    rune
		side     Side
		expected uint64
		err      error
	}{
		{
			move:     "e2e4",
			piece:    'p',
			side:     White,
			expected: board.White.Pawn - imath.Exp2(12) + imath.Exp2(28),
			err:      nil,
		}, {
			move:     "e7f5",
			piece:    'p',
			side:     Black,
			expected: board.Black.Pawn - imath.Exp2(52) + imath.Exp2(37),
			err:      nil,
		}, {
			move:     "a3e5",
			piece:    'q',
			side:     White,
			expected: board.White.Queen,
			err:      ErrNoPieceFound,
		}, {
			move:     "b1c3",
			piece:    'n',
			side:     White,
			expected: board.White.Knight - imath.Exp2(1) + imath.Exp2(18),
			err:      nil,
		},
	}

	for _, c := range cases {
		pieces := board.GetSide(c.side)
		fmt.Printf("Before: %b\n", *(pieces.GetPiece(c.piece)))
		board.Print()
		err := movePiece(c.move, pieces)
		fmt.Printf("After: %b\n", *(pieces.GetPiece(c.piece)))
		board.Print()
		if err != nil && err != c.err {
			t.Errorf(
				"Fail: Incorrected error\nCase: %s %q %s\nExpected: %v\nActual: %v",
				c.move, c.piece, c.side,
				c.err, err,
			)
		}

		if *(pieces.GetPiece(c.piece)) != c.expected {
			t.Errorf(
				"Fail: Result != Expected\nCase: %s %q %s\nExpected: %b\nActual: %b",
				c.move, c.piece, c.side,
				c.expected, *(pieces.GetPiece(c.piece)),
			)
		}
	}
}

func TestUnMovePiece(t *testing.T) {
	var board Board
	board.SetNew()

	cases := []struct {
		move     string
		piece    rune
		side     Side
		capture  uint8
		expected uint64
	}{}

	for _, c := range cases {
		fmt.Print(c)
	}
}
