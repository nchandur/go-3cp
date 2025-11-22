package main

import (
	"fmt"

	"github.com/nchandur/go-3cp/models"
)

func main() {

	outcomes := make([]int, 7)

	deck := models.NewDeck()

	for range 100 {
		player := models.NewHand(deck)
		outcome := player.Detect()
		outcomes[outcome]++
		deck.Reset()

	}

	fmt.Println(outcomes)

}
