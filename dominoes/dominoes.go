package dominoes

import "fmt"

// Define the Domino type here.
type Domino [2]int

func MakeChain(input []Domino) ([]Domino, bool) {
	fmt.Println(input)
	return input, false
}
