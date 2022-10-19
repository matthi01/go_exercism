package stringset

import "fmt"

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

// Define the Set type here.

type Set map[string]struct{}

func New() Set {
	return make(Set)
}

func NewFromSlice(l []string) Set {
	newSet := New()
	for _, v := range l {
		newSet.Add(v)
	}
	return newSet
}

func (s Set) String() string {
	var setLength int = len(s)
	var out string
	if len(s) == 0 {
		return "{}"
	}
	out += fmt.Sprintf("{")
	for key := range s {
		out += fmt.Sprintf("\"%s\"", key)
		if setLength-1 > 0 {
			out += fmt.Sprintf(", ")
		}
		setLength--
	}
	out += fmt.Sprintf("}")
	return out
}

func (s Set) IsEmpty() bool {
	if len(s) == 0 {
		return true
	}
	return false
}

func (s Set) Has(elem string) bool {
	if _, ok := s[elem]; ok {
		return true
	}
	return false
}

func (s Set) Add(elem string) {
	s[elem] = struct{}{}
}

func Subset(s1, s2 Set) bool {
	if len(s1) > len(s2) {
		return false
	}
	for k1 := range s1 {
		if _, ok := s2[k1]; !ok {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	return false
}

func Equal(s1, s2 Set) bool {
	fmt.Println("###")
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	if len(s1) != len(s2) {
		return false
	}
	for k1 := range s1 {
		for k2 := range s2 {
			if k1 != k2 {
				return false
			}
		}
	}
	return true
}

func Intersection(s1, s2 Set) Set {
	intersectionSet := New()
	for k1 := range s1 {
		if _, ok := s2[k1]; ok {
			intersectionSet.Add(k1)
		}
	}
	return intersectionSet
}

func Difference(s1, s2 Set) Set {
	differenceSet := New()
	for k1 := range s1 {
		if _, ok := s2[k1]; !ok {
			differenceSet.Add(k1)
		}
	}
	return differenceSet
}

func Union(s1, s2 Set) Set {
	unionSet := New()
	for k1 := range s1 {
		unionSet.Add(k1)
	}
	for k2 := range s2 {
		unionSet.Add(k2)
	}
	return unionSet
}
