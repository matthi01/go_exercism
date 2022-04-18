package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	// check if unit is correct
	if _, ok := units[unit]; !ok {
		return false
	}
	// check duplicate item
	if _, ok := bill[item]; !ok {
		bill[item] = units[unit]
	} else {
		bill[item] += units[unit]
	}
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	if _, ok := units[unit]; !ok {
		return false
	}
	if _, ok := bill[item]; !ok {
		return false
	}
	numberOnBill := bill[item]
	numberToRemove := units[unit]
	if numberOnBill-numberToRemove < 0 {
		return false
	} else if numberOnBill-numberToRemove == 0 {
		delete(bill, item)
		return true
	} else {
		bill[item] -= numberToRemove
		return true
	}
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	if _, ok := bill[item]; !ok {
		return 0, false
	}
	return bill[item], true
}
