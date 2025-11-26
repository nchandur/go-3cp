package main

import (
	"log"
	"os"

	"github.com/nchandur/go-3cp/game"
)

func main() {

	f, err := os.OpenFile("files/dev.logs", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	game := game.NewGame()

	game.Play()

}
