package bottlesong

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

type Translations map[int]string

var nTranslations Translations = Translations{
	0:  "no",
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
}

func Recite(startBottles, takeDown int) []string {
	if startBottles < 0 || startBottles > 10 {
		return []string{}
	}
	output := []string{}
	nBottles := startBottles
	for i := takeDown; i > 0; i-- {
		introVerse, err := getIntroVerse(nBottles)
		if err != nil {
			log.Fatal(err)
		}
		output = append(output, introVerse...)

		output = append(output, getActionVerse())

		resultVerse, err := getResultVerse(nBottles - 1)
		if err != nil {
			log.Fatal(err)
		}
		output = append(output, resultVerse)

		if i > 1 {
			output = append(output, "")
		}

		nBottles--
	}
	return output
}

func getIntroVerse(n int) ([]string, error) {
	if n < 0 {
		return []string{}, errors.New("invalid input provided")
	}
	if _, ok := nTranslations[n]; !ok {
		return []string{}, errors.New("invalid input provided")
	}
	nTranslation := nTranslations[n]
	bottleText := "bottles"
	if n == 1 {
		bottleText = "bottle"
	}
	output := []string{}
	for i := 0; i < 2; i++ {
		output = append(output,
			fmt.Sprintf("%s green %s hanging on the wall,", strings.Title(nTranslation), bottleText),
		)
	}
	return output, nil
}

func getActionVerse() string {
	return fmt.Sprintf("And if one green bottle should accidentally fall,")
}

func getResultVerse(n int) (string, error) {
	if n < 0 {
		return "", errors.New("invalid input provided")
	}
	if _, ok := nTranslations[n]; !ok {
		return "", errors.New("invalid input provided")
	}
	nTranslation := nTranslations[n]
	bottleText := "bottles"
	if n == 1 {
		bottleText = "bottle"
	}
	return fmt.Sprintf("There'll be %s green %s hanging on the wall.", nTranslation, bottleText), nil
}
