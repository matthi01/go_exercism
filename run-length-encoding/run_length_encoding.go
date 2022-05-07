package encode

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func RunLengthEncode(input string) string {
	var encoded string
	if input == "" {
		return encoded
	}
	var currentChar rune
	var currentCharCount int
	for _, c := range input {
		if currentChar == c {
			currentCharCount++
		} else {
			if currentChar != 0 {
				var strCount = fmt.Sprintf("%d", currentCharCount)
				if currentCharCount == 1 {
					strCount = ""
				}
				encoded += fmt.Sprintf("%s%s", strCount, string(currentChar))
			}
			currentChar = c
			currentCharCount = 1
		}
	}
	var strCount = fmt.Sprintf("%d", currentCharCount)
	if currentCharCount == 1 {
		strCount = ""
	}
	encoded += fmt.Sprintf("%s%s", strCount, string(currentChar))
	return encoded
}

func RunLengthDecode(input string) string {
	var decoded string
	if input == "" {
		return decoded
	}
	var reComponents = regexp.MustCompile("\\d+\\D?|\\D?")
	var reSplitComponent = regexp.MustCompile("\\d+|\\D?")
	// each 'component' contains the count and the character in a single string
	// each 'subComponent' is an array of either one or two items, first the count, second
	// the character - if the count is omitted the character will be used exactly once
	components := reComponents.FindAllString(input, -1)
	for _, component := range components {
		subComponents := reSplitComponent.FindAllString(component, -1)
		if len(subComponents) == 1 {
			decoded += subComponents[0]
		}
		if len(subComponents) == 2 {
			count, err := strconv.Atoi(subComponents[0])
			if err != nil {
				continue
			}
			character := subComponents[1]
			decoded += fmt.Sprintf("%s", strings.Repeat(character, count))
		}
	}
	return decoded
}
