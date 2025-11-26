package dealer

import (
	"fmt"
	"strings"

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

func (d *Dealer) Play() error {
	fmt.Printf("Dealer Hand: %s\n%s\n", strings.ToTitle(models.HandMap[d.Detect()]), d.String())
	return nil
}
