// Package weather contains weather releated functions.
package weather

// CurrentCondition indicates the current weather condition.
var CurrentCondition string

// CurrentLocation indicates the current location.
var CurrentLocation string

// Forecast provides the weather conditions for given location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
