package intmath

import (
	"math"
	"testing"
)

func TestBasicExponential(t *testing.T) {
	cases := []struct {
		base int
		exp  int
	}{
		{
			base: 2,
			exp:  2,
		},
		{
			base: 10,
			exp:  4,
		},
	}

	for _, c := range cases {
		intExp := Pow(c.base, c.exp)
		fltExp := math.Pow(float64(c.base), float64(c.exp))

		if float64(intExp) != fltExp {
			t.Errorf(
				"Fail case %d^%d:\nExpected: %f\nActual: %d\n",
				c.base, c.exp,
				fltExp, intExp,
			)
		}
	}
}

func TestExp2(t *testing.T) {
	cases := []struct {
		input    uint64
		expected uint64
	}{
		{
			input:    0,
			expected: 1,
		}, {
			input:    7,
			expected: 128,
		}, {
			input:    1,
			expected: 2,
		}, {
			input:    32,
			expected: 4294967296,
		},
	}

	for _, c := range cases {
		result := Exp2(c.input)
		if result != c.expected {
			t.Errorf(
				"Fail: Case: %d\nExpected: %d\nActual: %d\n",
				c.input, c.expected, result,
			)
		}
	}
}

func TestGetIndices(t *testing.T) {
	cases := []struct {
		input    uint64
		expected []int
	}{
		{
			input:    1,
			expected: []int{0},
		}, {
			input:    4294967296,
			expected: []int{32},
		}, {
			input:    10,
			expected: []int{1, 3},
		}, {
			input:    0b010100011101011010110101011010101,
			expected: []int{0, 2, 4, 6, 7, 9, 11, 13, 14, 16, 18, 19, 21, 23, 24, 25, 29, 31},
		},
	}

	for _, c := range cases {
		result := GetIndices(c.input)
		if len(result) != len(c.expected) {
			t.Errorf(
				"Fail: case %d\nExpected: %v\nActual: %v\n",
				c.input,
				c.expected,
				result,
			)
		}

		for i, val := range result {
			if val != c.expected[i] {
				t.Errorf(
					"Fail: case %d at index %d\nExpected: %v\nActual: %v\n",
					c.input, i,
					c.expected,
					result,
				)
			}
		}
	}
}
