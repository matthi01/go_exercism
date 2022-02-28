// Package twofer is short for two for one. One for you and one for me. It provides a name interpolation function to receive the expected phrase.
package twofer

import "fmt"

// ShareWith provides name interpolation to receive the expected phrase.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
