package dealer

import (
	"fmt"

	"github.com/nchandur/go-3cp/models"
)

type Dealer struct {
	*models.Hand
	*models.Stats
	Kaasu int
}

func NewDealer(deck *models.Deck) *Dealer {
	hand := models.NewHand(deck)
	stats := models.NewStats()
	return &Dealer{Hand: hand, Stats: stats}
}

func (d *Dealer) Play() (string, error) {
	fmt.Printf("Dealer Hand: \n%s\n%s\n", d.String(), models.HandMap[d.Detect()])

	// disqualified
	if d.Detect() == 0 && d.Hand.Cards[2].GetValue() < 11 {
		return "disqualified", nil
	}

	return "", nil
}
