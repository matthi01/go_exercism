package robotname

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

// idea:
// generate a map of sequential names
// assign one after another

// Define the Robot type here.
type Robot struct {
	name string
}

var usedNames = make(map[string]bool)

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numberOfLetters = 2
const numberOfNumbers = 3

var maxValidNames = int(math.Pow(float64(len(letters)), numberOfLetters) * math.Pow(10.0, numberOfNumbers))

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		r.Reset()
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = createUniqueName()
}

func createUniqueName() string {
	var name string = ""
	var valid bool = false
	var err error = nil
	for !valid && err == nil {
		name = createRandomLetters(numberOfLetters) + createRandomNumberString(numberOfNumbers)
		valid, err = validName(name)
	}
	usedNames[name] = true
	return name
}

func createRandomLetters(n int) string {
	var l []byte
	for i := 0; i < n; i++ {
		j := getRandomNumber(len(letters))
		l = append(l, letters[j])
	}
	return string(l)
}

func createRandomNumberString(n int) string {
	var ns string
	for i := 0; i < n; i++ {
		ns = ns + fmt.Sprintf("%d", getRandomNumber(10))
	}
	return ns
}

func getRandomNumber(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

func validName(name string) (bool, error) {
	fmt.Println("max valid names:", maxValidNames)
	fmt.Println("names allocated:", len(usedNames))
	if len(usedNames) == maxValidNames {
		return false, errors.New("out of valid names")
	}
	if _, exists := usedNames[name]; exists {
		return false, nil
	}
	return true, nil
}
