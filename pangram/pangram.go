package pangram

import (
	"strings"
)

var alphabet string = "abcdefghijklmnopqrstuvwxyz"

func IsPangram(input string) bool {
	if len(input) < len(alphabet) {
		return false
	}
	inputMap := make(map[rune]bool)
	input = strings.ToLower(input)
	for _, char := range input {
		if _, ok := inputMap[char]; ok {
			continue
		}
		inputMap[char] = true
	}
	for _, char := range alphabet {
		if _, ok := inputMap[char]; !ok {
			return false
		}
	}
	return true
}
