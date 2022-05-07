package prime

import "math"

func Factors(i int64) []int64 {
	var factors = []int64{}
	for i%2 == 0 {
		factors = append(factors, 2)
		i /= 2
	}
	for j := int64(3); j <= int64(math.Sqrt(float64(i))); j += 2 {
		for i%j == 0 {
			factors = append(factors, j)
			i /= j
		}
	}
	if i >= 3 {
		factors = append(factors, i)
	}
	return factors
}
