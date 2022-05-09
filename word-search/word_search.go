package wordsearch

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	var results = map[string][2][2]int{}
	if len(puzzle) == 0 {
		return results, errors.New("no puzzle provided")
	}
	height := len(puzzle)
	width := utf8.RuneCountInString(puzzle[0])
	if height != width {
		return results, errors.New("unevem puzzle grid provided")
	}
	for _, line := range puzzle {
		if utf8.RuneCountInString(line) != width {
			return results, errors.New("uneven line lengths detected in provided puzzle grid")
		}
	}
	for _, word := range words {
		fmt.Println(word)
	}
	return map[string][2][2]int{}, nil
}
