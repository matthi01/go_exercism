// Package accumulate accumulates stuff
package accumulate

// Accumulate takes a collection and does a specified operation on that collection
func Accumulate(ss []string, f func(string) string) []string {
	var result []string
	for _, s := range ss {
		result = append(result, f(s))
	}
	return result
}
