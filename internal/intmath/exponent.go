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

func Exp2(exp uint64) (result uint64) {
	result = 1 << exp
	return

}
