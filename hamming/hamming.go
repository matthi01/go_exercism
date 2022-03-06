// Package hamming calculates the hamming distance of DNA strands
package hamming

import "errors"

// Distance calculates the hamming distance of two DNA strands
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("input DNA strands must be of the same length")
	}
	var count int
	for i, char := range a {
		if char != rune(b[i]) {
			count++
		}
	}
	return count, nil
}
