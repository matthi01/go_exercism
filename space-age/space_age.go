package space

import (
	"math"
)

type Planet string

type OrbitalPeriod struct {
	planet     Planet
	earthYears float64
}

var secondsToEarthYear = 31557600.0

var orbitalPeriods = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1.0,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

func Age(age float64, planet Planet) float64 {
	if _, ok := orbitalPeriods[planet]; !ok {
		return 0.0
	}
	orbitalPeriod := orbitalPeriods[planet]
	localAge := (age / secondsToEarthYear) / orbitalPeriod
	return math.Round(localAge*100) / 100
}
