package armstrong

import (
	"math"
	"strconv"
	"unicode/utf8"
)

func IsNumber(n int) bool {
	strN := strconv.Itoa(n)
	exp := utf8.RuneCountInString(strN)
	result := 0.0
	for _, s := range strN {
		currentDigit, err := strconv.Atoi(string(s))
		if err != nil {
			return false
		}
		result += math.Pow(float64(currentDigit), float64(exp))
	}

	return result == float64(n)
}
