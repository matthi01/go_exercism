package change

import (
	"errors"
	"fmt"
	"sort"
)

func Change(coins []int, target int) ([]int, error) {
	var result = []int{}
	if len(coins) == 0 {
		return result, errors.New("no coins provided")
	}
	if target < 1 {
		return result, errors.New("no sensible target provided")
	}

	variation := 3
	for i := 0; i < variation; i++ {
		curCombination := []int{}
		remainingAmount := target
		highestCoin := 0
		for remainingAmount > 0 {
			test, _ := findDenominatorList(coins, remainingAmount)
			fmt.Println("list: ", test)
			applicableCoins, ok := findDenominatorList(coins, remainingAmount)
			if len(applicableCoins) > i {
				highestCoin = applicableCoins[i]
			} else {
				highestCoin = applicableCoins[0]
			}
			if !ok {
				return []int{}, errors.New("amount cannot be broken up by given coins")
			}
			remainingAmount -= highestCoin
			curCombination = append(curCombination, highestCoin)
			fmt.Println(curCombination)
		}
		if len(result) == 0 || len(curCombination) < len(result) {
			result = curCombination
		}
	}

	return result, nil
}

func findDenominatorList(list []int, value int) ([]int, bool) {
	denominatorList := []int{}
	if len(list) == 0 || value < 1 {
		return denominatorList, false
	}
	for _, v := range list {
		if value%v == 0 {
			denominatorList = append(denominatorList, v)
		}
	}
	sort.Slice(denominatorList, func(i, j int) bool {
		return denominatorList[i] > denominatorList[j]
	})
	return denominatorList, true
}

// func findHighestDenominator(list []int, value int) (int, bool) {
// 	if len(list) == 0 || value < 1 {
// 		return 0, false
// 	}
// 	var highest int = 0
// 	var multipliers int = 0
// 	var denominatorFound bool
// 	for _, v := range list {
// 		if value%v != 0 {
// 			continue
// 		}
// 		m := value / v
// 		if multipliers == 0 || m < multipliers {
// 			multipliers = m
// 			highest = v
// 			denominatorFound = true
// 		}
// 	}
// 	return highest, denominatorFound
// }
