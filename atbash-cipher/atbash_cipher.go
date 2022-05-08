package atbash

var alphabet = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

// var encodedAlphabet = []rune{'','','','','','','','','','','','','','','','','','','','','','','d','c','b','a',}

func Atbash(s string) string {
	// fmt.Println("#####")
	// fmt.Println("input:", s)
	// fmt.Println("alphabet:", alphabet)
	// fmt.Println("reverse:", reverseRuneSlice(alphabet))
	// fmt.Println("alphabet:", alphabet)

	var encoded string

	var reversedAlphabet = reverseRuneSlice(alphabet)

	for _, r := range s {
		if index, ok := findIndex(alphabet, r); ok {
			encodedRune := reversedAlphabet[index]
			encoded += string(encodedRune)
			continue
		}
		encoded += string(r)
	}
	return encoded
}

func reverseRuneSlice(sr []rune) []rune {
	var newSr = make([]rune, len(sr))
	copy(newSr, sr)
	for i, j := 0, len(newSr)-1; i < j; i, j = i+1, j-1 {
		newSr[i], newSr[j] = newSr[j], newSr[i]
	}
	return newSr
}

func findIndex(sr []rune, value rune) (int, bool) {
	var index int
	var ok bool
	for i, r := range sr {
		if r == value {
			index = i
			ok = true
			break
		}
	}
	return index, ok
}
