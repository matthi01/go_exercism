package logs

import (
	"fmt"
	"unicode/utf8"
)

var applicationMap = map[rune]string{
	10071:  "recommendation",
	128269: "search",
	9728:   "weather",
}

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, r := range log {
		if a, ok := applicationMap[r]; ok {
			return a
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var newLog string
	for _, r := range log {
		if r == oldRune {
			newLog += fmt.Sprintf("%c", newRune)
			continue
		}
		newLog += fmt.Sprintf("%c", r)
	}
	return newLog
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
