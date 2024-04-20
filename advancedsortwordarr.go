package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := range a {
		for j := i + 1; j < len(a); j++ {
			if f(string(a[i]), string(a[j])) == 1 {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
