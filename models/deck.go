package models

import (
	"math/rand/v2"
)

type Deck []Card

func NewDeck() *Deck {

	deck := make(Deck, 0)

	ranks := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	values := []uint8{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

	for i, rank := range ranks {
		for _, suit := range []string{"C", "H", "S", "D"} {
			card := Card{value: values[i], rank: rank, suit: suit}
			deck = append(deck, card)
		}

	}

	return &deck

}

func (d *Deck) Draw() *Card {

	idx := rand.IntN(len(*d))

	card := (*d)[idx]

	(*d) = append((*d)[:idx], (*d)[idx+1:]...)

	return &card
}
