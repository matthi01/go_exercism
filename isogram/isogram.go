package isogram

import (
	"strings"
)

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	word = strings.ReplaceAll(word, " ", "")
	word = strings.ReplaceAll(word, "-", "")
	letters := make(map[rune]bool)
	for _, l := range strings.ToLower(word) {
		if _, ok := letters[l]; ok {
			return false
		}
		letters[l] = true
	}
	return true
}
