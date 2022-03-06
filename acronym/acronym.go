// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import "strings"

// Abbreviate creates an acronym from the provided string
func Abbreviate(s string) string {
	s1 := strings.ReplaceAll(s, "-", " ")
	s2 := strings.ReplaceAll(s1, "_", " ")
	ss := strings.Split(s2, " ")
	var acr string
	for _, word := range ss {
		if len(word) > 0 {
			acr += string(word[0])
		}
	}
	acr = strings.ToUpper(acr)
	return acr
}
