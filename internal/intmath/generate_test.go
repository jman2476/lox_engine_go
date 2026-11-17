package intmath

import (
	"testing"
)

func TestGenHorz(t *testing.T) {
	cases := []struct {
		input    uint64
		expected []uint64
	}{
		{
			input:    6,
			expected: []uint64{0, 1, 2, 3, 4, 5, 6, 7},
		}, {
			input:    63,
			expected: []uint64{56, 57, 58, 59, 60, 61, 62, 63},
		}, {
			input:    33,
			expected: []uint64{32, 33, 34, 35, 36, 37, 38, 39},
		}, {
			input:    18,
			expected: []uint64{16, 17, 18, 19, 20, 21, 22, 23},
		}, {
			input:    67,
			expected: []uint64{},
		},
	}

	for _, c := range cases {
		result := GenHorizontal(c.input)
		if len(c.expected) != len(result) {
			t.Errorf(
				"Fail: incorrect result length\nExpected: %v %d\nActual: %v %d\n",
				c.expected, len(c.expected),
				result, len(result),
			)
			continue
		}

		for i, r := range result {
			if r != c.expected[i] {
				t.Errorf(
					"Fail: unexpected result\nCase: %d\nExpected: %v\nActual: %v\n",
					c.input, c.expected, result,
				)
				break
			}
		}

	}
}

func TestGenVert(t *testing.T) {
	cases := []struct {
		input    uint64
		expected []uint64
	}{
		{
			input:    6,
			expected: []uint64{6, 14, 22, 30, 38, 46, 54, 62},
		}, {
			input:    63,
			expected: []uint64{7, 15, 23, 31, 39, 47, 55, 63},
		}, {
			input:    33,
			expected: []uint64{1, 9, 17, 25, 33, 41, 49, 57},
		}, {
			input:    18,
			expected: []uint64{2, 10, 18, 26, 34, 42, 50, 58},
		}, {
			input:    67,
			expected: []uint64{},
		},
	}

	for _, c := range cases {
		result := GenVertical(c.input)
		if len(c.expected) != len(result) {
			t.Errorf(
				"Fail: incorrect result length\nExpected: %v %d\nActual: %v %d\n",
				c.expected, len(c.expected),
				result, len(result),
			)
			continue
		}

		for i, r := range result {
			if r != c.expected[i] {
				t.Errorf(
					"Fail: unexpected result\nCase: %d\nExpected: %v\nActual: %v\n",
					c.input, c.expected, result,
				)
				break
			}
		}

	}
}

func TestGenForeDiag(t *testing.T) {
	cases := []struct {
		input    uint64
		expected []uint64
	}{
		{
			input:    34,
			expected: []uint64{16, 25, 34, 43, 52, 61},
		}, {
			input:    4,
			expected: []uint64{4, 13, 22, 31},
		}, {
			input:    7,
			expected: []uint64{7},
		}, {
			input:    29,
			expected: []uint64{2, 11, 20, 29, 38, 47},
		},
	}

	for _, c := range cases {
		result := GenForeDiag(c.input)
		if len(c.expected) != len(result) {
			t.Errorf(
				"Fail: incorrect result length\nExpected: %v %d\nActual: %v %d\n",
				c.expected, len(c.expected),
				result, len(result),
			)
			continue
		}

		for i, r := range result {
			if r != c.expected[i] {
				t.Errorf(
					"Fail: unexpected result\nCase: %d\nExpected: %v\nActual: %v\n",
					c.input, c.expected, result,
				)
				break
			}
		}
	}
}

func TestGenBackDiag(t *testing.T) {
	cases := []struct {
		input    uint64
		expected []uint64
	}{
		{
			input:    49,
			expected: []uint64{7, 14, 21, 28, 35, 42, 49, 56},
		}, {
			input:    0,
			expected: []uint64{0},
		}, {
			input:    37,
			expected: []uint64{23, 30, 37, 44, 51, 58},
		}, {
			input:    13,
			expected: []uint64{6, 13, 20, 27, 34, 41, 48},
		}, {
			input:    47,
			expected: []uint64{47, 54, 61},
		}, {
			input:    8,
			expected: []uint64{1, 8},
		}, {
			input:    6,
			expected: []uint64{6, 13, 20, 27, 34, 41, 48},
		}, {
			input:    48,
			expected: []uint64{6, 13, 20, 27, 34, 41, 48},
		},
	}

	for _, c := range cases {
		result := GenBackDiag(c.input)
		if len(c.expected) != len(result) {
			t.Errorf(
				"Fail: incorrect result length\nCase: %d\nExpected: %v %d\nActual: %v %d\n",
				c.input,
				c.expected, len(c.expected),
				result, len(result),
			)
			continue
		}

		for i, r := range result {
			if r != c.expected[i] {
				t.Errorf(
					"Fail: unexpected result\nCase: %d\nExpected: %v\nActual: %v\n",
					c.input, c.expected, result,
				)
				break
			}
		}
	}
}

func TestOnAFile(t *testing.T) {
	cases := []struct {
		input    uint64
		expected bool
	}{
		{
			input:    54,
			expected: false,
		}, {
			input:    16,
			expected: true,
		}, {
			input:    56,
			expected: true,
		}, {
			input:    55,
			expected: false,
		},
	}

	for _, c := range cases {
		result := OnAFile(c.input)
		if c.expected != result {
			t.Errorf(
				"Fail: case %d\nExpected: %t\nActual: %t",
				c.input, c.expected, result,
			)
		}
	}
}

func TestOnHFile(t *testing.T) {
	cases := []struct {
		input    uint64
		expected bool
	}{
		{
			input:    54,
			expected: false,
		}, {
			input:    47,
			expected: true,
		}, {
			input:    56,
			expected: false,
		}, {
			input:    55,
			expected: true,
		},
	}

	for _, c := range cases {
		result := OnHFile(c.input)
		if c.expected != result {
			t.Errorf(
				"Fail: case %d\nExpected: %t\nActual: %t",
				c.input, c.expected, result,
			)
		}
	}
}
