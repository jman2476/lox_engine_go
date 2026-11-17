package intmath

import "testing"

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
		result := genHorizontal(c.input)
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
					"Fail: unexpected result\nExpected: %v\nActual: %v\n",
					c.expected, result,
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
		result := genVertical(c.input)
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
					"Fail: unexpected result\nExpected: %v\nActual: %v\n",
					c.expected, result,
				)
				break
			}
		}

	}
}
