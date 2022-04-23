package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	cows int
}

var DivisionByZeroError error = errors.New("division by zero")
var NegativeFodderError error = errors.New("negative fodder")

func (e SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	if cows < 0 {
		return 0.0, SillyNephewError{
			cows: cows,
		}
	}
	if cows == 0 {
		return 0.0, DivisionByZeroError
	}

	// double fodderAmount if we encounter a scale malfunction error
	fodderAmount, err := weightFodder.FodderAmount()
	if err == ErrScaleMalfunction {
		fodderAmount *= 2
	} else if err != nil {
		return 0.0, err
	}

	if fodderAmount < 0 {
		return 0.0, NegativeFodderError
	}

	amountPerCow := fodderAmount / float64(cows)
	return amountPerCow, nil
}
