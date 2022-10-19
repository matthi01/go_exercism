package sublist

import (
	"reflect"
)

type Relation string

func Sublist(l1, l2 []int) Relation {
	if len(l1) == len(l2) {
		for i, v := range l1 {
			if l2[i] != v {
				return "unequal"
			}
		}
		return "equal"
	}
	if isSublist(l1, l2) == true {
		return "sublist"
	}
	if isSublist(l2, l1) == true {
		return "superlist"
	}
	return "unequal"
}

func isSublist(l1, l2 []int) bool {
	if len(l1) > len(l2) {
		return false
	}
	if len(l1) == 0 {
		return true
	}
	for _, v1 := range l1 {
		for i2, v2 := range l2 {
			if len(l2) < i2+len(l1) {
				break
			}
			if v1 == v2 {
				if reflect.DeepEqual(l1, l2[i2:i2+len(l1)]) {
					return true
				}
			}
		}
	}
	return false
}
