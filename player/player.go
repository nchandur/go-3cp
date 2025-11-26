package player

import (
	"fmt"
	"strings"

	"github.com/nchandur/go-3cp/models"
)

type Player struct {
	*models.Hand
	*models.Stats
	*models.Payout
	Kaasu int
}

func NewPlayer(deck *models.Deck) *Player {
	hand := models.NewHand(deck)
	stats := models.NewStats()
	payout := models.Payout{}
	return &Player{Hand: hand, Stats: stats, Payout: &payout}
}

func (p *Player) Play() (string, error) {

	fmt.Printf("\n\nPlayer hand: %s\n%s\n", models.HandMap[p.Detect()], p.Hand.String())

	var play string
	fmt.Printf("Continue playing? (y/n): ")

	_, err := fmt.Scanln(&play)

	if err != nil {
		return "", fmt.Errorf("failed to play: %v", err)
	}

	play = strings.ToLower(play)

	if play[0] == 'y' {
		return "", nil
	}

	if play[0] == 'n' {
		return "quit", nil
	}

	return "", fmt.Errorf("invalid input: must be y/n")
}
