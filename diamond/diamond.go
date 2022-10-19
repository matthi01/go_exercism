package diamond

import (
	"errors"
	"fmt"
	"strings"
)

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

const (
	desc = "descending"
	asc  = "ascending"
)

func Gen(char byte) (string, error) {
	if char < 'A' || char > 'Z' {
		return "", errors.New("Input out of range")
	}
	var depth int
	for i, c := range alphabet {
		if c == rune(char) {
			depth = i
		}
	}

	var diamond string
	var newLine string
	for level := 0; level <= depth; level++ {
		newLine = addNewLine(level, depth, asc)
		diamond = diamond + newLine
	}
	for level := depth - 1; level >= 0; level-- {
		newLine = addNewLine(level, depth, desc)
		diamond = diamond + newLine
	}
	return diamond, nil
}

func addNewLine(level, depth int, direction string) string {
	var newLine string
	if level == 0 {
		// single
		formatStr := "%s%c%s\n"
		if direction == desc || depth == 0 {
			formatStr = "%s%c%s"
		}
		newLine = fmt.Sprintf(formatStr, strings.Repeat(" ", depth), alphabet[level], strings.Repeat(" ", depth))
	} else {
		// double
		innerOffset := level
		if level > 1 {
			innerOffset += (level - 1)
		}
		newLine = fmt.Sprintf("%s%c%s%c%s\n", strings.Repeat(" ", depth-level), alphabet[level], strings.Repeat(" ", innerOffset), alphabet[level], strings.Repeat(" ", depth-level))
	}
	return newLine
}
