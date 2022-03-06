// Package perfect contains number classification logic
package perfect

import (
	"errors"
)

// Classification represents the number classification type
type Classification int

// Classification types...
const (
	ClassificationPerfect   Classification = 1
	ClassificationDeficient Classification = 2
	ClassificationAbundant  Classification = 3
)

// ErrOnlyPositive is an error indicating that only positive numbers are allowed
var ErrOnlyPositive error = errors.New("only positive numbers")

// ErrUnknown is an error indicating something unknown has happened
var ErrUnknown error = errors.New("unknown error has been encountered")

// Classify will return the classification of a given number
func Classify(n int64) (Classification, error) {
	var factorSum int64
	if n < 1 {
		return 0, ErrOnlyPositive
	}
	var i int64
	for i = 1; i < n; i++ {
		if n%i == 0 {
			factorSum += i
		}
	}
	switch {
	case factorSum == n:
		return ClassificationPerfect, nil
	case factorSum < n:
		return ClassificationDeficient, nil
	case factorSum > n:
		return ClassificationAbundant, nil
	default:
		return 0, ErrUnknown
	}
}
