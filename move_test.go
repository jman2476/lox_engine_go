package main

import "testing"

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
		expected int
		err      error
	}{
		{},
	}
}
