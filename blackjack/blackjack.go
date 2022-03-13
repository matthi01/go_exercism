// Package blackjack is a representation of the game blackjack
package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	deck := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"ten":   10,
		"jack":  10,
		"queen": 10,
		"king":  10,
		"ace":   11,
	}
	if val, ok := deck[card]; ok {
		return val
	}
	return 0
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	points := ParseCard(card1) + ParseCard(card2)
	return points == 21
}

// - Stand (S)
// - Hit (H)
// - Split (P)
// - Automatically win (W)

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	if isBlackjack {
		if dealerScore == 10 || dealerScore == 11 {
			return "S"
		}
		return "W"
	}
	return "P"
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
	switch {
	case handScore >= 17:
		return "S"
	case handScore <= 11:
		return "H"
	case handScore > 11 && handScore < 17:
		if dealerScore >= 7 {
			return "H"
		}
		return "S"
	default:
		return ""
	}
}
