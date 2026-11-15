package main

import (
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
		result, err := findPiece(c.input, pieces)
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
		},
	}

	for _, c := range cases {
		pieces := board.GetSide(c.side)
		err := movePiece(c.move, &pieces)
		if err != nil && err != c.err {
			t.Errorf(
				"Fail: Incorrected error\nCase: %s %q %s\nExpected: %v\nActual: %v",
				c.move, c.piece, c.side,
				c.err, err,
			)
		}

		if pieces.GetPiece(c.piece) != c.expected {
			t.Errorf(
				"Fail: Result != Expected\nCase: %s %q %s\nExpected: %b\nActual: %b",
				c.move, c.piece, c.side,
				c.expected, pieces.GetPiece(c.piece),
			)
		}
	}
}
