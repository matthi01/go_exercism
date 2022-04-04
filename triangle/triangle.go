package triangle

import (
	"sort"
)

// Kind indicates the type of triangle
type Kind int

const (
	// NaT - not a triangle
	NaT = 0
	// Equ - equilateral
	Equ = 1
	// Iso - isosceles
	Iso = 2
	// Sca - scalene
	Sca = 3
)

// KindFromSides determines the type of triangle given the length of its sides
func KindFromSides(a, b, c float64) Kind {
	// check for negative inputs
	if a <= 0 || b <= 0 || c <= 0 {
		return 0
	}
	// check for triangle inequality & degenerate triangles
	sortedSides := []float64{a, b, c}
	sort.Float64s(sortedSides)
	if sortedSides[0]+sortedSides[1] <= sortedSides[2] {
		return 0
	}

	if a == b && b == c {
		return 1
	}
	if a == b || a == c || b == c {
		return 2
	}
	if a != b && b != c {
		return 3
	}
	return 0
}
