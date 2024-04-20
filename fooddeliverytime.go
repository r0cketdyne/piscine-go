package piscine

// Define a struct for each menu item
type food struct {
	preptime int
}

// Define a map to store the menu items and their corresponding preparation times
var menu = map[string]food{
	"burger":  {preptime: 15},
	"chips":   {preptime: 10},
	"nuggets": {preptime: 12},
}

// Function to calculate the total preparation time for the order
func FoodDeliveryTime(order string) int {
	totalTime := 0
	food, ok := menu[order]
	if !ok {
		return 404
	}
	totalTime = food.preptime
	return totalTime
}
