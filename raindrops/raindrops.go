// Package raindrops seems to be a version of fizzbuzz
package raindrops

import "fmt"

// Convert seems to be a version of fizzbuzz
func Convert(number int) string {
	var str string
	var factorFound bool
	if number%3 == 0 {
		str += "Pling"
		factorFound = true
	}
	if number%5 == 0 {
		str += "Plang"
		factorFound = true
	}
	if number%7 == 0 {
		str += "Plong"
		factorFound = true
	}
	if !factorFound {
		str = fmt.Sprint(number)
	}
	return str
}
