package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	var re, err = regexp.Compile(`^\[TRC\]|^\[DBG\]|^\[INF\]|^\[WRN\]|^\[ERR\]|^\[FTL\]\w+`)
	if err != nil {
		return false
	}
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	var re, err = regexp.Compile(`<[-=~*]*>`)
	if err != nil {
		return []string{}
	}
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	var re, err = regexp.Compile(`(?i)".*password.*"`)
	if err != nil {
		return 0
	}
	var counter int
	for _, line := range lines {
		if re.MatchString(line) {
			counter++
		}
	}
	return counter
}

func RemoveEndOfLineText(text string) string {
	var re, err = regexp.Compile(`end-of-line\d*`)
	if err != nil {
		return ""
	}
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	var re, err = regexp.Compile(`User\s+(\w+)`)
	if err != nil {
		return []string{}
	}
	var updatedLines = []string{}
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if matches != nil {
			line = fmt.Sprintf("[USR] %s %s", matches[1], line)
		}
		updatedLines = append(updatedLines, line)
	}
	return updatedLines
}
