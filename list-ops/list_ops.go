package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	var acc = initial
	for _, n := range s {
		acc = fn(acc, n)
	}
	return acc
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	var acc = initial
	for i := len(s) - 1; i >= 0; i-- {
		acc = fn(s[i], acc)
	}
	return acc
}

func (s IntList) Filter(fn func(int) bool) IntList {
	var filtered = IntList{}
	for _, n := range s {
		if fn(n) {
			filtered = append(filtered, n)
		}
	}
	return filtered
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	var mappedList = IntList{}
	for _, n := range s {
		mappedList = append(mappedList, fn(n))
	}
	return mappedList
}

func (s IntList) Reverse() IntList {
	var reversedList = IntList{}
	for i := len(s) - 1; i >= 0; i-- {
		reversedList = append(reversedList, s[i])
	}
	return reversedList
}

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	var newList = make(IntList, len(s))
	copy(newList, s)
	for _, l := range lists {
		newList = newList.Append(l)
	}
	return newList
}
