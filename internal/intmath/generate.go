package intmath

// Determine squares on the same:
//		- horizontal
//		- vertical
//		- back diagonal
//		- fore diagonal

func genHorizontal(square uint64) (horz []uint64) {
	if square > 63 {
		return
	}

	aEdge := square - (square % 8)
	hEdge := aEdge + 7

	for i := aEdge; i <= hEdge; i++ {
		horz = append(horz, i)
	}

	return
}
