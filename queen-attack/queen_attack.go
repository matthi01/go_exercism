package queenattack

import (
	"errors"
	"math"
)

var (
	ErrSameSquare    = errors.New("same square")
	ErrInvalidSquare = errors.New("invalid square")
)

// files == columns == character == first position
// ranks == rows == number == second position
var validFiles = []rune{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H',
}
var validRanks = []rune{
	'1', '2', '3', '4', '5', '6', '7', '8',
}

type square struct {
	file rune
	rank rune
}

func CanQueenAttack(w string, b string) (bool, error) {
	if w == b {
		return false, ErrSameSquare
	}
	squares := []string{w, b}
	for _, s := range squares {
		err := validateSquare(s)
		if err != nil {
			return false, err
		}
	}

	// for readability... and playing around
	white := square{
		file: rune(w[0]),
		rank: rune(w[1]),
	}
	black := square{
		file: rune(b[0]),
		rank: rune(b[1]),
	}

	// same file / same rank
	if (white.file == black.file) || (white.rank == black.rank) {
		return true, nil
	}
	// common diagonals
	if math.Abs(float64(white.file-black.file)) == math.Abs(float64(white.rank-black.rank)) {
		return true, nil
	}

	// no attack
	return false, nil
}

func validateSquare(s string) error {
	if len(s) != 2 {
		return ErrInvalidSquare
	}
	if !contains(validFiles, rune(s[0])) || !contains(validRanks, rune(s[1])) {
		return ErrInvalidSquare
	}
	return nil
}

func contains(cx []rune, r rune) bool {
	for _, c := range cx {
		if c == r {
			return true
		}
	}
	return false
}
