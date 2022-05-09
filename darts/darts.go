package darts

import (
	"math"
)

var scoreMap = map[string]int{
	"outside": 0,
	"outer":   1,
	"middle":  5,
	"inner":   10,
}

func Score(x, y float64) int {
	var score = 0
	r := math.Sqrt(x*x + y*y)
	switch {
	case r <= 1.0:
		score += scoreMap["inner"]
	case r <= 5.0:
		score += scoreMap["middle"]
	case r <= 10.0:
		score += scoreMap["outer"]
	default:
		score += scoreMap["outside"]
	}
	return score
}
