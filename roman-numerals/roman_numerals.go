package romannumerals

import "errors"

var i1 = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
var i2 = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
var i3 = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
var i4 = []string{"", "M", "MM", "MMM"}

func ToRomanNumeral(input int) (string, error) {
	var roman string
	if input <= 0 || input > 3000 {
		return "", errors.New("invalid range")
	}

	// getting number at each decimal place
	roman += i4[(input%1e4)/1e3]
	roman += i3[(input%1e3)/1e2]
	roman += i2[(input%1e2)/1e1]
	roman += i1[(input % 1e1)]

	return roman, nil
}
