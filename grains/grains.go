package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("only positive numbers between 1 and 64 are allowed")
	}
	if number == 1 {
		return 1, nil
	}
	return uint64(math.Pow(2, float64(number-1))), nil
}

func Total() uint64 {
	var total uint64
	for i := 1; i <= 64; i++ {
		grains, err := Square(i)
		if err != nil {
			return 0
		}
		total += grains
	}
	return total
}
