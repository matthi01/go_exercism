package etl

import (
	"strings"
)

func Transform(in map[int][]string) map[string]int {
	var outputMap = make(map[string]int)
	for pointValue, charSlice := range in {
		for _, char := range charSlice {
			char = strings.ToLower(char)
			outputMap[char] += pointValue
		}
	}
	return outputMap
}
