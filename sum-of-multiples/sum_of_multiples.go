package summultiples

func SumMultiples(limit int, divisors ...int) int {
	var sum int
	var multiples = make(map[int]bool)
	for i := 1; i < limit; i++ {
		for _, divisor := range divisors {
			if divisor == 0 {
				continue
			}
			if i%divisor == 0 {
				if _, ok := multiples[i]; !ok {
					multiples[i] = true
				}
			}
		}
	}
	for n := range multiples {
		sum += n
	}
	return sum
}
