// Package diffsquares needs a description
package diffsquares

import (
	"math"
)

// SquareOfSum sums the factors of a natural number and squares it
func SquareOfSum(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}
	return int(math.Pow(float64(sum), 2))
}

// SumOfSquares sums the square of each factor of a natural number and sums them
func SumOfSquares(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += int(math.Pow(float64(i), 2))
	}
	return sum
}

// Difference finds the difference between the SumOfSquares and SquareOfSum for a given number
func Difference(n int) int {
	return int(math.Abs(float64(SumOfSquares(n) - SquareOfSum(n))))
}
