// Package cars contains analytical tools for a car production line
package cars

import (
	"math"
)

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(math.Floor(float64(productionRate/60) * (successRate / 100)))
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	sets := carsCount / 10
	individualCount := carsCount % 10

	var totalCost uint
	totalCost += uint(sets * 95000)
	totalCost += uint(individualCount * 10000)
	return totalCost
}
