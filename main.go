package main

import (
	"fmt"

	"github.com/nchandur/go-3cp/models"
	"github.com/nchandur/go-3cp/player"
)

func main() {

	deck := models.NewDeck()
	player := player.NewPlayer(deck)

	fmt.Println(player)

}
