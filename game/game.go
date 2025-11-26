package game

import (
	"fmt"

	"github.com/nchandur/go-3cp/dealer"
	"github.com/nchandur/go-3cp/models"
	"github.com/nchandur/go-3cp/player"
)

var anteBonusPayouts map[int]uint64 = map[int]uint64{
	0: 0, // high card
	1: 0, // one pair
	2: 0, // flush
	3: 1, // straight
	4: 4, // three of a kind
	5: 5, // straight flush
	6: 6, // royal flush
}

var bonusPayouts map[int]uint64 = map[int]uint64{
	0: 0,
	1: 1,
	2: 4,
	3: 6,
	4: 30,
	5: 40,
	6: 80,
}

type Game struct {
	*models.Deck
	*player.Player
	*dealer.Dealer
}

func NewGame() *Game {
	game := Game{
		Deck: models.NewDeck(),
	}
	p := player.NewPlayer(game.Deck)
	d := dealer.NewDealer(game.Deck)

	game.Player = p
	game.Dealer = d

	return &game
}

func (g *Game) Play() error {

	for {
		var ante uint64
		var bonus uint64

		fmt.Printf("Place ante: ")
		_, err := fmt.Scanln(&ante)

		if err != nil {
			return fmt.Errorf("failed to play game: %v", err)
		}

		fmt.Printf("Place bonus: ")
		_, err = fmt.Scanln(&bonus)

		if err != nil {
			return fmt.Errorf("failed to play game: %v", err)
		}

		playOut, err := g.Player.Play()

		if err != nil {
			return fmt.Errorf("failed to play game: %v", err)
		}

		if playOut == "quit" {
			break
		}

		err = g.Dealer.Play()

		if err != nil {
			return fmt.Errorf("failed to play game: %v", err)
		}

		g.Player.Hand = models.NewHand(g.Deck)
		g.Dealer.Hand = models.NewHand(g.Deck)

		g.Deck = models.NewDeck()

		g.Payouts(ante, bonus)
		fmt.Println(g.Player.Payout)
	}

	return nil
}

func (g *Game) Payouts(ante, bonus uint64) {

	var playPay = func() {

		// check for dealer qualification
		if g.Dealer.Hand.Cards[2].GetValue() < 12 && g.Dealer.Hand.Detect() == 0 {
			g.Player.Payout.Ante += ante
			return
		}

		compared := g.Player.Hand.Compare(g.Dealer.Hand)

		switch compared {

		// victory
		case 1:
			g.Player.Payout.Ante += (ante + (ante * anteBonusPayouts[g.Player.Detect()]))
			g.Player.Payout.Play += ante

		// loss
		case -1:
			g.Player.Payout.Ante -= ante
			g.Player.Payout.Play -= ante
		}

	}

	var bonusPay = func() {
		g.Player.Payout.Bonus += (bonus + (bonus * bonusPayouts[g.Player.Detect()]))
	}

	// calculates payouts for ante and play wagers
	playPay()

	// calculates pair plus bonus payouts
	bonusPay()

}
