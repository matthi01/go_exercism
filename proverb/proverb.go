package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	var output []string
	lastIndex := len(rhyme) - 1
	for i, str := range rhyme {
		if i == lastIndex {
			output = append(output, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
			break
		}
		output = append(output, fmt.Sprintf("For want of a %s the %s was lost.", str, rhyme[i+1]))
	}
	return output
}
