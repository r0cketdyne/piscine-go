package piscine

func UltimateDivMod(a *int, b *int) {
	a1 := *a
	b1 := *b
	div1 := a1 / b1
	mod1 := a1 % b1
	*a = div1
	*b = mod1
}
