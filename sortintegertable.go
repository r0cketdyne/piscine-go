package piscine

func SortIntegerTable(table []int) {
	for i := 0; i < len(table); i++ {
		for j := len(table) - 1; j > i; j-- {
			if table[i] > table[j] {
				temp := table[i]
				table[i] = table[j]
				table[j] = temp
			}
		}
	}
}
