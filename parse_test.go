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
