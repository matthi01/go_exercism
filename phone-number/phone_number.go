package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

var (
	ErrInvalidInput        = errors.New("invalid  input")
	ErrInvalidAreaCode     = errors.New("invalid area code")
	ErrInvalidExchangeCode = errors.New("invalid exchange code")
	ErrInvalidPhoneNumber  = errors.New("invalid phone number")
)

func Number(phoneNumber string) (string, error) {
	re := regexp.MustCompile("[0-9]+")
	sn := re.FindAllString(phoneNumber, -1)
	if len(sn) == 0 {
		return "", ErrInvalidInput
	}
	parsedNumber := ""
	for _, s := range sn {
		parsedNumber += s
	}
	parsedNumber = strings.ReplaceAll(parsedNumber, " ", "")

	// if 11 digits, first digit must be a one to be safely dropped
	numLength := utf8.RuneCountInString(parsedNumber)
	if numLength < 10 || numLength > 11 {
		return "", ErrInvalidInput
	}
	if numLength == 11 {
		if parsedNumber[:1] != "1" {
			return "", ErrInvalidInput
		}
		parsedNumber = parsedNumber[1:]
	}

	areaCode := parsedNumber[:3]
	exchangeCode := parsedNumber[3:6]

	// area code cannot start with a 0 or 1
	if areaCode[0] == '0' || areaCode[0] == '1' {
		return "", ErrInvalidAreaCode
	}

	// exchange code cannot start with a 0 or 1
	if exchangeCode[0] == '0' || exchangeCode[0] == '1' {
		return "", ErrInvalidExchangeCode
	}
	return parsedNumber, nil
}

func AreaCode(phoneNumber string) (string, error) {
	parsedNumber, err := Number(phoneNumber)
	if err != nil {
		return "", ErrInvalidPhoneNumber
	}
	areaCode := parsedNumber[0:3]
	return areaCode, nil
}

func Format(phoneNumber string) (string, error) {
	parsedNumber, err := Number(phoneNumber)
	if err != nil {
		return "", ErrInvalidPhoneNumber
	}
	areaCode := parsedNumber[:3]
	exchangeCode := parsedNumber[3:6]
	subscriberCode := parsedNumber[6:]
	return fmt.Sprintf("(%s) %s-%s", areaCode, exchangeCode, subscriberCode), nil
}
