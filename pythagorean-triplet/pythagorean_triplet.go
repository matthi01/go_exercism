package pythagorean

import (
	"fmt"
	"math"
)

type Triplet [3]int

func Range(min, max int) []Triplet {
	var triplets []Triplet
	// maybe put some thought into it.. there should be nicer way
	for a := min; a <= max; a++ {
		aSq := math.Pow(float64(a), 2)
		for b := a; b <= max; b++ {
			bSq := math.Pow(float64(b), 2)
			for c := b; c <= max; c++ {
				cSq := math.Pow(float64(c), 2)
				if aSq+bSq == cSq && a < b && b < c {
					triplets = append(triplets, Triplet{a, b, c})
				}
			}
		}
	}
	return triplets
}

func Sum(p int) []Triplet {
	fmt.Println("start...")
	var triplets []Triplet
	var options = Range(0, p)
	fmt.Println("starting option loop...")
	for _, triplet := range options {
		if triplet[0]+triplet[1]+triplet[2] == p {
			triplets = append(triplets, triplet)
		}
	}
	fmt.Println("done...")
	return triplets
}
