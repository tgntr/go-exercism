package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := make(map[string]int, 6)
	units["quarter_of_a_dozen"] = 3
	units["half_of_a_dozen"] = 6
	units["dozen"] = 12
	units["small_gross"] = 120
	units["gross"] = 144
	units["great_gross"] = 1728

	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int, 6)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	value, exists := units[unit]
	if !exists {
		return false
	} else {
		bill[item] += value
		return true
	}
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {

	if billValue, billExists := bill[item]; !billExists {
		return false
	} else if unitValue, unitExists := units[unit]; !unitExists {
		return false
	} else if billValue < unitValue {
		return false
	} else if billValue == unitValue {
		delete(bill, item)
		return true
	} else {
		bill[item] -= unitValue
		return true
	}

}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (value int, exists bool) {
	value, exists = bill[item]
	return
}
