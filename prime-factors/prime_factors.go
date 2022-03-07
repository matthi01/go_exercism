// Package prime provides prime factor calculation functions
package prime

// Factors calculates the prime factors of a given number
func Factors(n int64) []int64 {
	remainder := n
	var factors []int64
	var counter int64 = 2
	for remainder >= counter {
		if remainder%counter == 0 {
			factors = append(factors, counter)
			remainder = remainder / counter
			continue
		}
		counter = nextPrime(counter)
	}
	return factors
}

func isPrime(n int64) bool {
	if n < 2 {
		return false
	}
	var i int64
	for i = 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func nextPrime(n int64) int64 {
	var next int64 = n + 1
	for {
		if isPrime(next) {
			return next
		}
		next++
	}
}
