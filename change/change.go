package change

import (
	"errors"
	"fmt"
)

func Change(coins []int, target int) ([]int, error) {
	fmt.Println("coins:", coins)
	fmt.Println("target:", target)
	var result = []int{}
	if len(coins) == 0 {
		return result, errors.New("no coins provided")
	}
	if target < 1 {
		return result, errors.New("no sensible target provided")
	}

	remainingAmount := target
	for remainingAmount > 0 {
		highestCoin, ok := findHighestDenominator(coins, remainingAmount)
		if !ok {
			return []int{}, errors.New("amount cannot be broken up by given coins")
		}
		remainingAmount -= highestCoin
		result = append(result, highestCoin)
	}

	return result, nil
}

func findHighestDenominator(list []int, value int) (int, bool) {
	if len(list) == 0 || value < 1 {
		return 0, false
	}
	var highest int = 0
	var multipliers int = 0
	var denominatorFound bool
	for _, v := range list {
		if value%v != 0 {
			continue
		}
		m := value / v
		if multipliers == 0 || m < multipliers {
			multipliers = m
			highest = v
			denominatorFound = true
		}
	}
	return highest, denominatorFound
}
