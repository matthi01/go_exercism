package prime

func Nth(n int) (int, bool) {
	var primes = []int{}
	if n < 1 {
		return 0, false
	}
	for i := 2; len(primes) < n; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes[len(primes)-1], true
}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
