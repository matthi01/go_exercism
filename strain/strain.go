package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	var filtered Ints
	for _, n := range i {
		if filter(n) {
			filtered = append(filtered, n)
		}
	}
	return filtered
}

func (i Ints) Discard(filter func(int) bool) Ints {
	var filtered Ints
	for _, n := range i {
		if !filter(n) {
			filtered = append(filtered, n)
		}
	}
	return filtered
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	var filtered Lists
	for _, li := range l {
		if filter(li) {
			filtered = append(filtered, li)
		}
	}
	return filtered
}

func (s Strings) Keep(filter func(string) bool) Strings {
	var filtered Strings
	for _, ss := range s {
		if filter(ss) {
			filtered = append(filtered, ss)
		}
	}
	return filtered
}
