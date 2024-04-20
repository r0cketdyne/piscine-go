package piscine

func DivMod(a int, b int, div *int, mod *int) {
	div1 := a / b
	mod1 := a % b
	*div = div1
	*mod = mod1
}
