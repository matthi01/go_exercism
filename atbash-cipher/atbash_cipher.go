package atbash

import (
	"strings"
)

const (
	exclusions = "#$%^&,!. "
	alphabet   = "abcdefghijklmnopqrstuvwxyz"
	groupSize  = 5
)

func Atbash(s string) string {
	var cleanedString string
	for _, r := range s {
		if strings.Contains(exclusions, string(r)) {
			continue
		}
		cleanedString += string(r)
	}
	cleanedString = strings.ToLower(cleanedString)

	var encryptedString string
	for _, r := range cleanedString {
		index := findIndex(string(r), string(alphabet))
		if index > -1 {
			encryptedString += string(alphabet[len(alphabet)-1-index])
			continue
		}
		encryptedString += string(r)
	}

	var grouped string
	for i, r := range encryptedString {
		if i != 0 && i%groupSize == 0 {
			grouped += " "
		}
		grouped += string(r)
	}
	return grouped
}

func findIndex(item, searchString string) int {
	for i, r := range searchString {
		if item == string(r) {
			return i
		}
	}
	return -1
}
