package sieve

import "sort"

func Sieve(limit int) []int {
	var numbers = make(map[int]bool, limit)
	var primes = []int{}
	for i := 2; i <= limit; i++ {
		// if found, means that we already marked it as a non-prime
		if _, ok := numbers[i]; ok {
			continue
		}
		// any unmarked number still available is a prime
		numbers[i] = true
		var multiple = i
		for multiple < limit {
			multiple = multiple + i
			numbers[multiple] = false
		}
	}
	for n, isPrime := range numbers {
		if isPrime {
			primes = append(primes, n)
		}
	}
	sort.Ints(primes)
	return primes
}
