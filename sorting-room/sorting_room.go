package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	switch fnb.(type) {
	case FancyNumber:
		n, err := strconv.Atoi(fnb.Value())
		if err != nil {
			return 0
		}
		return n
	default:
		return 0
	}
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	n := ExtractFancyNumber(fnb)
	return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(n))
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	var strReturn string
	switch i.(type) {
	case float64:
		strReturn = DescribeNumber(i.(float64))
	case int:
		strReturn = DescribeNumber(float64(i.(int)))
	case NumberBox:
		strReturn = DescribeNumberBox(i.(NumberBox))
	case FancyNumberBox:
		strReturn = DescribeFancyNumberBox(i.(FancyNumberBox))
	default:
		strReturn = "Return to sender"
	}
	return strReturn
}
