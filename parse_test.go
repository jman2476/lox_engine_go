package main

import "testing"

func TestParseSquare(t *testing.T) {
	cases := []struct {
		square      string
		expected    int
		expectedErr error
	}{
		{
			square:   "a1",
			expected: 0,
		},
		{
			square:   "c5",
			expected: 34,
		},
		{
			square:      "b0",
			expected:    -1,
			expectedErr: ErrInvalidSquare,
		},
		{
			square:      "g9",
			expected:    -1,
			expectedErr: ErrInvalidSquare,
		}, {
			square:   "d8",
			expected: 59,
		},
	}

	for _, c := range cases {
		result, resErr := parseSquare(c.square)
		if c.expected != result {
			t.Errorf(
				"Fail: case %s\nExpected: %d\nActual: %d\n",
				c.square,
				c.expected,
				result,
			)
		}

		if c.expectedErr != resErr {
			t.Errorf(
				"Fail: case %s\nExpected: %v\nActual: %v\n",
				c.square,
				c.expectedErr,
				resErr,
			)
		}
	}
}

func TestDigCharToInt(t *testing.T) {
	cases := []struct {
		input    rune
		expected int
	}{
		{
			input:    '1',
			expected: 1,
		},
		{
			input:    '5',
			expected: 5,
		},
		{
			input:    'a',
			expected: 49,
		},
		{
			input:    '9',
			expected: 9,
		},
	}

	for _, c := range cases {
		result := digitCharToInt(c.input)
		if result != c.expected {
			t.Errorf(
				"Fail: case %q\nExpected: %d\nActual: %d",
				c.input,
				c.expected,
				result,
			)
		}
	}
}

func TestFileCharToInt(t *testing.T) {
	cases := []struct {
		input    rune
		expected int
	}{
		{
			input:    'b',
			expected: 1,
		},
		{
			input:    'f',
			expected: 5,
		},
		{
			input:    'a',
			expected: 0,
		},
		{
			input:    'i',
			expected: 8,
		},
	}

	for _, c := range cases {
		result := fileCharToInt(c.input)
		if result != c.expected {
			t.Errorf(
				"Fail: case %q\nExpected: %d\nActual: %d",
				c.input,
				c.expected,
				result,
			)
		}
	}
}

func TestParseMove(t *testing.T) {
	cases := []struct {
		input string
		from  int
		to    int
		promo rune
		err   error
	}{
		{
			input: "a1b5",
			from:  0,
			to:    33,
			promo: 0,
		},
		{
			input: "a1b5q",
			from:  0,
			to:    33,
			promo: 113,
		},
		{
			input: "h4c4",
			from:  31,
			to:    26,
			promo: 0,
		},
		{
			input: "g7f8n",
			from:  54,
			to:    61,
			promo: 110,
		},
		{
			input: "a0b5",
			from:  -1,
			to:    -1,
			promo: 0,
			err:   ErrInvalidSquare,
		},
	}

	for _, c := range cases {
		from, to, promo, err := parseMove(c.input)
		if from != c.from {
			t.Errorf(
				"Fail: case %s\nExpected: %v\nActual: %v\n",
				c.input,
				c.from,
				from,
			)
		}
		if to != c.to {
			t.Errorf(
				"Fail: case %s\nExpected: %v\nActual: %v\n",
				c.input,
				c.to,
				to,
			)
		}
		if promo != c.promo {
			t.Errorf(
				"Fail: case %s\nExpected: %v\nActual: %v\n",
				c.input,
				c.promo,
				promo,
			)
		}
		if err != nil && err != c.err {
			t.Errorf(
				"Fail: case %s\nExpected: %v\nActual: %v\n",
				c.input,
				c.err,
				err,
			)
		}
	}
}
