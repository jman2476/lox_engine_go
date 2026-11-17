package intmath

import (
	"slices"
)

// Determine squares on the same:
//		- horizontal
//		- vertical
//		- back diagonal
//		- fore diagonal

func GenHorizontal(square uint64) (horz []uint64) {
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

func GenVertical(square uint64) (vert []uint64) {
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

func GenForeDiag(square uint64) (diag []uint64) {
	if square > 63 {
		return
	}

	for i := square; i < 64; i += 9 {
		diag = append(diag, i)
		if i%8 == 7 {
			break
		}
	}

	if square%8 != 0 && square > 8 {
		for i := square - 9; i < 64; i -= 9 {
			diag = append(diag, i)
			if i%8 == 0 {
				break
			}
		}
	}

	slices.Sort(diag)
	return
}

func GenBackDiag(square uint64) (diag []uint64) {
	if square > 63 {
		return
	}

	for i := square; i < 64; i += 7 {
		diag = append(diag, i)
		if i%8 == 0 {
			break
		}
	}

	if square%8 != 7 && square > 7 {
		for i := square - 7; i < 64; i -= 7 {
			diag = append(diag, i)
			if i%8 == 7 {
				break
			}
		}
	}

	slices.Sort(diag)
	return
}
