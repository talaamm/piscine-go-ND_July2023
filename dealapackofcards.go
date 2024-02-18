package piscinego

import "fmt"

func DealAPackOfCards(deck []int) {
	cardIndex := 0
	for playerNum := 1; playerNum <= 4; playerNum += 1 {
		fmt.Printf("Player %v: %v, %v, %v\n", playerNum, deck[cardIndex], deck[cardIndex+1], deck[cardIndex+2])
		cardIndex += 3
	}
}
