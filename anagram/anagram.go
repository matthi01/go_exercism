package anagram

import (
	"reflect"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	var validCandidates []string
	subjectMap := mapifyWord(subject)
	for _, candidate := range candidates {
		if strings.ToLower(subject) == strings.ToLower(candidate) {
			continue
		}
		candidateMap := mapifyWord(candidate)
		isEqual := reflect.DeepEqual(subjectMap, candidateMap)
		if isEqual {
			validCandidates = append(validCandidates, candidate)
		}
	}
	return validCandidates
}

func mapifyWord(word string) map[rune]int {
	chars := make(map[rune]int)
	for _, char := range strings.ToLower(word) {
		if _, ok := chars[char]; !ok {
			chars[char] = 1
		} else {
			chars[char]++
		}
	}
	return chars
}
