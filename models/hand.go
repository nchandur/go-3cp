package models

type Hand []Card

func NewHand(deck *Deck) *Hand {
	hand := make(Hand, 0)

	for range 3 {
		hand = append(hand, *deck.Draw())
	}

	return &hand
}
