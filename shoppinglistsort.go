package piscine

func ShoppingListSort(slice []string) []string {
	for i := range slice {
		for j := i + 1; j < len(slice); j++ {
			if len(slice[i]) > len(slice[j]) {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	return slice
}
