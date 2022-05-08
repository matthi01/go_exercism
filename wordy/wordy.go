package wordy

import (
	"regexp"
	"strconv"
)

func Answer(question string) (int, bool) {
	var re = regexp.MustCompile("-?\\d+|plus|minus|multiplied by|divided by|cubed")
	components := re.FindAllString(question, -1)

	if len(components) == 1 {
		if n, err := strconv.Atoi(components[0]); err == nil {
			return n, true
		}
		return 0, false
	}

	if len(components)%2 == 0 {
		return 0, false
	}

	var result int
	var nextOperation string
	for i, component := range components {
		// special case if this is the first one
		if i == 0 {
			if n, err := strconv.Atoi(component); err == nil {
				result = n
				continue
			}
			return 0, false
		}

		// component is an operation (should be anyway)
		if i%2 != 0 {
			nextOperation = component
			continue
		} else {
			// component should be a number
			if n, err := strconv.Atoi(component); err == nil {
				switch nextOperation {
				case "plus":
					result += n
				case "minus":
					result -= n
				case "multiplied by":
					result *= n
				case "divided by":
					result /= n
				default:
					return 0, false
				}
				nextOperation = ""
				continue
			}
			// couldn't convert what we expected to be a number
			return 0, false
		}
	}
	return result, true
}
