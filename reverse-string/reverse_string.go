// Package reverse reverses a string
package reverse

// Reverse reverses a string
func Reverse(input string) string {
	str := []rune(input)
	var strNew []rune
	for i := len(str) - 1; i >= 0; i-- {
		strNew = append(strNew, str[i])
	}
	return string(strNew)
}
