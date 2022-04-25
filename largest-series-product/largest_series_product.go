package lsproduct

import (
	"errors"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span == 0 {
		return 1, nil
	}
	if span < 0 {
		return 0, errors.New("cannot have a negative span")
	}
	if span > len(digits) || len(digits) == 0 {
		return 0, errors.New("span cannot be larger than input")
	}
	var greatestProduct int64
	var currentProduct int64
	for i := 0; i <= len(digits)-span; i++ {
		series := digits[i : span+i]
		for i, c := range series {
			d, err := strconv.Atoi(string(c))
			if err != nil {
				return 0, errors.New("encountered conversion error - bad input")
			}
			if i == 0 {
				currentProduct = int64(d)
				continue
			}
			currentProduct *= int64(d)
		}
		if currentProduct > greatestProduct {
			greatestProduct = currentProduct
		}
	}
	return greatestProduct, nil
}
