package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

var re = regexp.MustCompile("[a-z0-9]+('[a-z0-9])*")

func WordCount(phrase string) Frequency {
	words := re.FindAllString(strings.ToLower(phrase), -1)
	instances := make(Frequency)
	for _, word := range words {
		instances[word]++
	}
	return instances
}
