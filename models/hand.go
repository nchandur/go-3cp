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

var HandMap map[int]string = map[int]string{
	HighCard:      "high card",
	OnePair:       "one pair",
	Flush:         "flush",
	Straight:      "straight",
	ThreeOfAKind:  "three of a kind",
	StraightFlush: "straight flush",
	RoyalFlush:    "royal flush",
}

type Hand struct {
	Cards []Card
	rank  string
}

func NewHand(deck *Deck) *Hand {
	hand := Hand{}

	for range 3 {
		hand.Cards = append(hand.Cards, *deck.Draw())
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

// hand that consists of two Cards of the same rank
func (h *Hand) isOnePair() bool {
	return (h.Cards[0].value == h.Cards[1].value) || (h.Cards[1].value == h.Cards[2].value)
}

// hand that consists of three Cards of the same suit, but that are not in consecutive ranking
func (h *Hand) isFlush() bool {
	return (h.Cards[0].suit == h.Cards[1].suit) && (h.Cards[0].suit == h.Cards[2].suit)
}

// hand that consists of three Cards that are in consecutive ranking, but that are not the same suit
func (h *Hand) isStraight() bool {
	return h.isConsecutive() && !h.isFlush()
}

// hand that consists of three Cards of the same rank
func (h *Hand) isThreeOfAKind() bool {
	return (h.Cards[0].value == h.Cards[1].value) && (h.Cards[0].value == h.Cards[2].value)
}

// hand that consists of three Cards of the same suit in consecutive ranking
func (h *Hand) isStraightFlush() bool {
	return h.isConsecutive() && h.isFlush()
}

// hand that consists of an ace, king, and queen of the same suit
func (h *Hand) isRoyalFlush() bool {
	return (h.Cards[0].rank == "Q" && h.Cards[1].rank == "K" && h.Cards[2].rank == "A") && h.isFlush()
}

// helper functions

func (h *Hand) String() string {

	var builder strings.Builder
	var cardLines [][]string

	for _, c := range (*h).Cards {
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

	sort.Slice(h.Cards, func(i, j int) bool {
		card1 := h.Cards[i]
		card2 := h.Cards[j]

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
	if h.Cards[2].rank == "A" {
		if ((h.Cards[0].value == 2) && (h.Cards[1].value == 3)) || ((h.Cards[0].value == 12) && (h.Cards[1].value == 13)) {
			return true
		}
	}

	// otherwise
	start := h.Cards[0].value
	return (h.Cards[1].value == (start + 1)) && (h.Cards[2].value == (start + 2))

}

func (h *Hand) HasRank(rank string) bool {
	for _, card := range (*h).Cards {
		if card.rank == rank {
			return true
		}
	}
	return false
}

func (h *Hand) Compare(hand *Hand) int8 {

	// better hand
	if h.Detect() > hand.Detect() {
		return 1
	}

	// worse hand
	if h.Detect() < hand.Detect() {
		return -1
	}

	// tie breaker
	for i := 2; i >= 0; i-- {

		if h.Cards[i].value > hand.Cards[i].value {
			return 1
		}
		if h.Cards[i].value < hand.Cards[i].value {
			return -1
		}

	}

	return 0

}
