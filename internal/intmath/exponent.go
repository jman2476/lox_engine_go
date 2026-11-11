package intmath

func Pow(base, exp int) int {
	result := 1
	for {
		if exp&1 == 1 {
			result *= base
		}
		exp >>= 1
		if exp == 0 {
			break
		}
		base *= base
	}
	return result
}

func Exp2(exp uint64) uint64 {
	return 1 << exp
}

func GetIndices(bitmap uint64) []int {
	var indices []int
	for i := range 64 {
		if bitmap == 0 {
			break
		}
		if bitmap&1 == 1 {
			indices = append(indices, i)
		}
		bitmap >>= 1
	}
	return indices
}
