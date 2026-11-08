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
