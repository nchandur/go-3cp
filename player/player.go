package player

import "github.com/nchandur/go-3cp/models"

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
