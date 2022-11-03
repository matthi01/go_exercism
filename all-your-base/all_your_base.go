package allyourbase

import (
	"errors"
	"math"
)

var (
	ErrInvalidInputBase  = errors.New("input base must be >= 2")
	ErrInvalidOutputBase = errors.New("output base must be >= 2")
	ErrInvalidDigits     = errors.New("all digits must satisfy 0 <= d < input base")
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return []int{0}, ErrInvalidInputBase
	}
	if outputBase < 2 {
		return []int{0}, ErrInvalidOutputBase
	}
	for _, d := range inputDigits {
		if d < 0 || d >= inputBase {
			return []int{0}, ErrInvalidDigits
		}
	}

	var outputDigits = []int{}
	var total int

	j := 0
	for i := len(inputDigits) - 1; i >= 0; i-- {
		d := inputDigits[j]
		total += d * int(math.Pow(float64(inputBase), float64(i)))
		j++
	}

	for total > 0 {
		r := total % outputBase
		total /= outputBase
		outputDigits = append(outputDigits, r)
	}

	outputDigits = reverse(outputDigits)

	if len(outputDigits) == 0 {
		return []int{0}, nil
	}

	return outputDigits, nil
}

func reverse(in []int) []int {
	var out []int = []int{}
	for i := len(in) - 1; i >= 0; i-- {
		out = append(out, in[i])
	}
	return out
}
