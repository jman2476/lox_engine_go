package intmath

import (
	"slices"
)

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

func genVertical(square uint64) (vert []uint64) {
	if square > 63 {
		return
	}

	for i := square; i < 64; i += 8 {
		vert = append(vert, i)
	}

	if square > 7 {
		for i := square - 8; i < 64; i -= 8 {
			vert = append(vert, i)
		}
	}

	slices.Sort(vert)
	return
}
