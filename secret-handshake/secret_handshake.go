package secret

import (
	"strconv"
	"strings"
)

func Handshake(code uint) []string {
	var handshake = []string{}
	var reverse bool = false
	if code < 1 {
		return handshake
	}
	binary := strconv.FormatUint(uint64(code), 2)
	digits := strings.Split(binary, "")

	var gestures = map[int]string{
		0: "wink",
		1: "double blink",
		2: "close your eyes",
		3: "jump",
	}

	for i := 0; i < len(gestures) && i < len(digits); i++ {
		if digits[(len(digits)-1)-i] == "1" {
			handshake = append(handshake, gestures[i])
		}
	}

	if len(digits) > 4 {
		if digits[4] == "1" {
			reverse = true
		}
	}

	if reverse {
		reverseSlice(handshake)
	}

	return handshake
}

func reverseSlice(ss []string) {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
}
