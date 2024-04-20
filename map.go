package piscine

func Map(f func(int) bool, a []int) []bool {
	b := []bool{}
	for _, i := range a {
		b = append(b, f(i))
	}
	return b
}
