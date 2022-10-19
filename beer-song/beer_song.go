package beer

import (
	"errors"
	"fmt"
)

func Song() string {
	song, _ := Verses(99, 0)
	return song
}

func Verses(start, stop int) (string, error) {
	if start < stop {
		return "", errors.New("stop cannot be ahead of start")
	}
	var output string
	for i := start; i >= stop; i-- {
		verse, err := Verse(i)
		if err != nil {
			return "", err
		}
		output += verse
		output += "\n"
	}
	return output, nil
}

func Verse(n int) (string, error) {
	if n < 0 || n > 99 {
		return "", errors.New("input must be an integer between 0 and 99")
	}
	var phrase string
	if n == 2 {
		phrase = fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottle of beer on the wall.\n", n, n, n-1)
	} else if n == 1 {
		phrase = fmt.Sprintf("%d bottle of beer on the wall, %d bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n", n, n)
	} else if n == 0 {
		phrase = fmt.Sprintf("No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n")
	} else {
		phrase = fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n", n, n, n-1)
	}

	return phrase, nil
}
