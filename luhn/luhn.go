// Package luhn uses the luhn alogrithm for number verification
package luhn

import (
	"strconv"
	"strings"
)

// Valid determines if a number is valid per the Luhn formula
func Valid(id string) bool {
	strID := strings.TrimSpace(id)
	strID = strings.ReplaceAll(strID, " ", "")
	if len(strID) <= 1 {
		return false
	}
	_, err := strconv.Atoi(strID)
	if err != nil {
		return false
	}

	var checkNumbers []int
	var staticNumbers []int
	// getting every second number starting from the right
	for i := len(strID) - 2; i >= 0; i -= 2 {
		n, err := strconv.Atoi(string(strID[i]))
		if err != nil {
			return false
		}
		checkNumbers = append(checkNumbers, n)
	}
	for i := len(strID) - 1; i >= 0; i -= 2 {
		n, err := strconv.Atoi(string(strID[i]))
		if err != nil {
			return false
		}
		staticNumbers = append(staticNumbers, n)
	}

	var sum int
	for _, n := range checkNumbers {
		n2 := n * 2
		if n2 > 9 {
			n2 -= 9
		}
		sum += n2
	}

	for _, n := range staticNumbers {
		sum += n
	}

	if sum%10 == 0 {
		return true
	}
	return false
}
