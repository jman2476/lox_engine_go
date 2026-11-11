package main

import (
	"errors"
	"fmt"
)

var (
	ErrMoveTooShort  = errors.New("incomplete move. Should be in format: [from square][to square][promotion piece] <all lower case>")
	ErrMoveTooLong   = errors.New("to many characters. Should be in format: [from square][to square][promotion piece] <all lower case>")
	ErrInvalidSquare = errors.New("square does not exist between a1 and h8")
)

func parseMove(move string) (from, to int, err error) {
	if len(move) < 4 {
		return -1, -1, ErrMoveTooShort
	}
	if len(move) > 5 {
		return -1, -1, ErrMoveTooLong
	}

	return 0, 0, nil
}

func parseSquare(square string) (int, error) {
	file := fileCharToInt(rune(square[0]))
	rank := digitCharToInt(rune(square[1]))

	fmt.Println(file, rank)
	if file < 0 || file > 7 {
		return -1, ErrInvalidSquare
	}
	if rank < 1 || rank > 8 {
		return -1, ErrInvalidSquare
	}
	return (rank-1)*8 + file, nil
}

func digitCharToInt(num rune) int {
	return int(num) - 48
}

func fileCharToInt(file rune) int {
	return int(file) - 97
}
