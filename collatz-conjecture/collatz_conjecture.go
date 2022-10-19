package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	var steps int
	if n < 1 {
		return 0, errors.New("input must be positive integer")
	}
	for {
		// emergency exit
		if steps > 10000 {
			return 0, errors.New("stuck in infinite loop")
		}
		if n == 1 {
			return steps, nil
		}
		// even
		if n%2 == 0 {
			n = n / 2
			steps++
		} else {
			n = (n * 3) + 1
			steps++
		}
	}
}
