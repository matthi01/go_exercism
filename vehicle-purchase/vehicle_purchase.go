// Package purchase contains helpful functions for purchasing a vehicle
package purchase

import (
	"fmt"
	"sort"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	options := []string{option1, option2}
	sort.Strings(options)
	return fmt.Sprintf("%s is clearly the better choice.", options[0])
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var value float64 = originalPrice
	if age < 3 {
		value = originalPrice * (float64(80) / 100)
	} else if age >= 3 && age < 10 {
		value = originalPrice * (float64(70) / 100)
	} else if age >= 10 {
		value = originalPrice * (float64(50) / 100)
	}
	return value
}
