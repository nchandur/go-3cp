package main

import (
	"fmt"
	"log"

	"github.com/nchandur/go-3cp/dealer"
	"github.com/nchandur/go-3cp/models"
	"github.com/nchandur/go-3cp/player"
)

func main() {

	deck := models.NewDeck()
	player := player.NewPlayer(deck)

	playerOutput, err := player.Play()

	if err != nil {
		log.Fatal(err)
	}

	if playerOutput == "quit" {
		return
	}

	dealer := dealer.NewDealer(deck)

	dealerOutput, err := dealer.Play()

	if dealerOutput == "disqualified" {
		fmt.Println("dealer disqualified")
	}

	fmt.Println(player.Hand.Compare(dealer.Hand))

}
