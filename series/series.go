package series

import "unicode/utf8"

func All(n int, s string) []string {
	var series = []string{}
	var inputLength = utf8.RuneCountInString(s)
	if n > inputLength {
		return series
	}
	var start = 0
	for start+n <= inputLength {
		series = append(series, s[start:start+n])
		start++
	}
	return series
}

func UnsafeFirst(n int, s string) string {
	allSubstrings := All(n, s)
	if len(allSubstrings) > 0 {
		return allSubstrings[0]
	}
	return ""
}

func First(n int, s string) (first string, ok bool) {
	allSubstrings := All(n, s)
	if len(allSubstrings) > 0 {
		return allSubstrings[0], true
	}
	return "", false
}
