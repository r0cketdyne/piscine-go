package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	count := 1
	playerCard1 := 0
	playerCard2 := 0

	for i := 1; i <= 12; i++ {
		if i%3 == 1 {
			playerCard1 = deck[i-1]
		} else if i%3 == 2 {
			playerCard2 = deck[i-1]
		}
		if i%3 == 0 {
			fmt.Printf("Player %d: %d, %d, %d\n", count, playerCard1, playerCard2, deck[i-1])
			count++
		}

	}
}
