// Package prime provides prime factor calculation functions
package prime

// Factors calculates the prime factors of a given number
func Factors(n int64) (factors []int64) {
	primes := findPrimes(n)
	remainder := n
	for i := 0; i < len(primes); {
		prime := primes[i]
		if remainder%prime == 0 {
			factors = append(factors, prime)
			remainder = remainder / prime
			continue
		}
		i++
	}
	return factors
}

func findPrimes(limit int64) (primes []int64) {
	arr := make([]bool, limit)
	var i int64
	var j int64
	if limit == 2 {
		return []int64{2}
	}
	for i = 2; i < limit; i++ {
		if arr[i] == true {
			continue
		}
		primes = append(primes, i)
		for j = i * i; j < limit; j += i {
			arr[j] = true
		}
	}
	return primes
}
