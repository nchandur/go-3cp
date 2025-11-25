package player

import (
	"fmt"
	"strings"

	"github.com/nchandur/go-3cp/models"
)

type Player struct {
	*models.Hand
	*models.Stats
	*models.Wager
	Kaasu int
}

func NewPlayer(deck *models.Deck) *Player {
	hand := models.NewHand(deck)
	stats := models.NewStats()
	return &Player{Hand: hand, Stats: stats}
}

func (p *Player) SetWagers(ante uint64, bonus uint64) {
	p.Wager = models.NewWager(ante, bonus)
}

func (p *Player) Play() (string, error) {
	var ante uint64
	var bonus uint64

	fmt.Printf("Place ante: ")
	_, err := fmt.Scanln(&ante)

	if err != nil {
		return "", fmt.Errorf("failed to play: %v", err)
	}

	fmt.Printf("Place bonus: ")
	_, err = fmt.Scanln(&bonus)

	if err != nil {
		return "", fmt.Errorf("failed to play: %v", err)
	}

	p.SetWagers(ante, bonus)

	fmt.Printf("\n\nPlayer hand: \n%s\n%s\n", p.Hand.String(), models.HandMap[p.Detect()])

	var play string
	fmt.Printf("Continue playing? (y/n): ")

	_, err = fmt.Scanln(&play)

	if err != nil {
		return "", fmt.Errorf("failed to play: %v", err)
	}

	play = strings.ToLower(play)

	if play[0] != 'n' {
		p.Wager.Play = p.Wager.Ante
		return "", nil
	}

	return "quit", nil
}
