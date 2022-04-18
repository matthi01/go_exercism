package isbn

import (
	"strconv"
	"strings"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) != 10 {
		return false
	}
	var checksum int
	for i, c := range isbn {
		var n int
		var err error
		if c == 'X' && i == len(isbn)-1 {
			n = 10
		} else {
			n, err = strconv.Atoi(string(c))
			if err != nil {
				return false
			}
		}
		checksum += n * (len(isbn) - i)
	}
	if checksum%11 == 0 {
		return true
	}
	return false
}
