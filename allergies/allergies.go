package allergies

var allergyMap = map[uint]string{
	1:   "eggs",
	2:   "peanuts",
	4:   "shellfish",
	8:   "strawberries",
	16:  "tomatoes",
	32:  "chocolate",
	64:  "pollen",
	128: "cats",
}

var allergyKeys = []uint{
	128, 64, 32, 16, 8, 4, 2, 1,
}

func Allergies(allergies uint) []string {
	var allergiesList = []string{}
	for allergies > 256 {
		allergies -= 256
	}
	for _, key := range allergyKeys {
		if allergies >= key {
			allergiesList = append(allergiesList, allergyMap[key])
			allergies -= key
		}
	}
	return allergiesList
}

func AllergicTo(allergies uint, allergen string) bool {
	allergiesList := Allergies(allergies)
	return contains(allergiesList, allergen)
}

func contains(ss []string, item string) bool {
	for _, s := range ss {
		if s == item {
			return true
		}
	}
	return false
}
