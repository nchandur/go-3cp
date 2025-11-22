package main

import (
	"fmt"

	"github.com/nchandur/go-3cp/models"
)

func main() {

	deck := models.NewDeck()
	player := models.NewHand(deck)

	fmt.Println(player)

}
