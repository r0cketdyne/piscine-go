package piscine

func ReverseMenuIndex(menu []string) []string {
	myMenu := make([]string, len(menu))
	for i := 0; i < len(menu); i++ {
		myMenu[len(menu)-1-i] = menu[i]
	}
	return myMenu
}
