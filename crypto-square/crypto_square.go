package cryptosquare

import (
	"fmt"
	"math"
	"strings"
)

var specialChars = " #$%^&,!."

func getDimensions(strLength int) (rows, columns int) {
	rows = int(math.Round(math.Sqrt(float64(strLength))))
	columns = int(math.Round(float64(strLength) / float64(rows)))
	if rows*columns < strLength {
		columns++
	}
	if rows*columns < strLength {
		rows++
	}
	return rows, columns
}

func Encode(pt string) string {
	var clean string
	for _, c := range pt {
		if strings.Contains(specialChars, string(c)) {
			continue
		}
		clean += string(c)
	}

	clean = strings.ToLower(clean)

	numRows, numColumns := getDimensions(len(clean))

	var rows []string
	var row string
	startIndex := 0
	for i := 0; i < numRows; i++ {
		if startIndex+numColumns > len(clean) {
			shortRow := clean[startIndex:]
			row = fmt.Sprintf("%s%s", shortRow, strings.Repeat(" ", numColumns-len(shortRow)))
		} else {
			row = clean[startIndex : startIndex+numColumns]
		}
		rows = append(rows, row)
		startIndex += numColumns
	}

	var encodedRows []string
	for column := 0; column < numColumns; column++ {
		var encodedRow string
		for _, row := range rows {
			encodedRow += string(row[column])
		}
		encodedRows = append(encodedRows, encodedRow)
	}

	encoded := strings.Join(encodedRows, " ")
	return encoded
}
