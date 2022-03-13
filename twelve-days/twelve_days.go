// Package twelve is probably the sumbest coding challenge I have ever done
package twelve

import "fmt"

var presents = []string{
	"Partridge in a Pear Tree",
	"Turtle Doves",
	"French Hens",
	"Calling Birds",
	"Gold Rings",
	"Geese-a-Laying",
	"Swans-a-Swimming",
	"Maids-a-Milking",
	"Ladies Dancing",
	"Lords-a-Leaping",
	"Pipers Piping",
	"Drummers Drumming",
}

var days = []string{
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

var presentNumbers = []string{
	"a",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
	"eleven",
	"twelve",
}

// Verse returns a single verse from the song
func Verse(n int) string {
	verse := fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s", days[n-1], presentList(n-1))
	if n == 1 {
		return verse
	}
	return fmt.Sprintf("%s", verse)
}

func presentList(n int) string {
	var l string
	var suffix string
	for i := n; i >= 0; i-- {
		if i == 0 {
			suffix = "."
		} else if i == 1 {
			suffix = ", and "
		} else {
			suffix = ", "
		}
		l += fmt.Sprintf("%s %s%s", presentNumbers[i], presents[i], suffix)
	}
	return l
}

// Song outputs the lyrics to this dumb song
func Song() string {
	var lyrics string
	for i := 1; i <= 12; i++ {
		lyrics += Verse(i)
		if i < 12 {
			lyrics += "\n"
		}
	}
	return lyrics
}
