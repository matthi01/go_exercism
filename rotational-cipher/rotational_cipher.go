package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	if shiftKey < 0 || shiftKey > 26 {
		return ""
	}
	output := []rune{}
	for _, char := range plain {
		if 'a' <= char && char <= 'z' {
			char = rune((byte(char)-byte('a')+byte(shiftKey))%26 + byte('a'))
		} else if 'A' <= char && char <= 'Z' {
			char = rune((byte(char)-byte('A')+byte(shiftKey))%26 + byte('A'))
		}
		output = append(output, char)
	}
	return string(output)
}
