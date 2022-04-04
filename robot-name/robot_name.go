package robotname

import (
	"errors"
	"fmt"
	"math"
)

type Robot struct {
	name string
}

var names = make(map[int]string)

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

const numberOfLetters = 2
const numberOfNumbers = 3

var maxValidNames = int(math.Pow(float64(len(letters)), numberOfLetters) * math.Pow(10.0, numberOfNumbers))
var currentNameIndex = 0

func init() {
	createNames()
}

func (r *Robot) Name() (string, error) {
	var err error
	if r.name == "" {
		err = r.Reset()
	}
	if err != nil {
		return "", err
	}
	return r.name, nil
}

func (r *Robot) Reset() error {
	var err error
	r.name, err = getNewName()
	if err != nil {
		return err
	}
	return nil
}

func getNewName() (string, error) {
	if currentNameIndex == maxValidNames {
		fmt.Println("out of names")
		return "", errors.New("out of valid names")
	}
	if _, ok := names[currentNameIndex]; !ok {
		return "", errors.New("couldn't retrieve name")
	}
	name := names[currentNameIndex]
	currentNameIndex++
	return name, nil
}

func createNames() {
	counter := 0
	for i := 0; i < len(letters); i++ {
		for j := 0; j < len(letters); j++ {
			for n := 0; n < 1000; n++ {
				name := fmt.Sprintf("%s%s%03d", string(letters[i]), string(letters[j]), n)
				names[counter] = name
				counter++
			}
		}
	}
}
