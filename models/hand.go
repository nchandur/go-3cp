package models

import (
	"sort"
	"strings"
)

const ()

const (
	HighCard = iota
	OnePair
	Flush
	Straight
	ThreeOfAKind
	StraightFlush
	RoyalFlush
)

type Hand struct {
	cards []Card
	rank  string
}

func NewHand(deck *Deck) *Hand {
	hand := Hand{}

	for range 3 {
		hand.cards = append(hand.cards, *deck.Draw())
	}

	hand.sort()

	return &hand
}

func (h *Hand) Detect() int {

	if h.isRoyalFlush() {
		return RoyalFlush
	}

	if h.isStraightFlush() {
		return StraightFlush
	}

	if h.isThreeOfAKind() {
		return ThreeOfAKind
	}

	if h.isStraight() {
		return Straight
	}

	if h.isFlush() {
		return Flush
	}

	if h.isOnePair() {
		return OnePair
	}
	return HighCard

}

// hand that consists of two cards of the same rank
func (h *Hand) isOnePair() bool {
	return (h.cards[0].value == h.cards[1].value) || (h.cards[1].value == h.cards[2].value)
}

// hand that consists of three cards of the same suit, but that are not in consecutive ranking
func (h *Hand) isFlush() bool {
	return (h.cards[0].suit == h.cards[1].suit) && (h.cards[0].suit == h.cards[2].suit)
}

// hand that consists of three cards that are in consecutive ranking, but that are not the same suit
func (h *Hand) isStraight() bool {
	return h.isConsecutive() && !h.isFlush()
}

// hand that consists of three cards of the same rank
func (h *Hand) isThreeOfAKind() bool {
	return (h.cards[0].value == h.cards[1].value) && (h.cards[0].value == h.cards[2].value)
}

// hand that consists of three cards of the same suit in consecutive ranking
func (h *Hand) isStraightFlush() bool {
	return h.isConsecutive() && h.isFlush()
}

// hand that consists of an ace, king, and queen of the same suit
func (h *Hand) isRoyalFlush() bool {
	return (h.cards[0].rank == "Q" && h.cards[1].rank == "K" && h.cards[2].rank == "A") && h.isFlush()
}

// helper functions

func (h *Hand) String() string {

	var builder strings.Builder
	var cardLines [][]string

	for _, c := range (*h).cards {
		cardLines = append(cardLines, c.String())
	}

	for i := range cardLines[0] {
		for _, card := range cardLines {
			builder.WriteString(card[i])
			builder.WriteString("  ")
		}
		builder.WriteString("\n")
	}

	return builder.String()

}

func (h *Hand) sort() {

	sort.Slice(h.cards, func(i, j int) bool {
		card1 := h.cards[i]
		card2 := h.cards[j]

		if card1.value != card2.value {
			return card1.value < card2.value
		}
		if card1.rank != card2.rank {
			return card1.rank < card2.rank
		}
		return card1.suit < card2.suit
	})

}

func (h *Hand) isConsecutive() bool {

	// sort hand just in case
	h.sort()

	// check ace cases
	if h.cards[2].rank == "A" {
		if ((h.cards[0].value == 2) && (h.cards[1].value == 3)) || ((h.cards[0].value == 12) && (h.cards[1].value == 13)) {
			return true
		}
	}

	// otherwise
	start := h.cards[0].value
	return (h.cards[1].value == (start + 1)) && (h.cards[2].value == (start + 2))

}
