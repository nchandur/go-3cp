package models

import "fmt"

type Payout struct {
	Ante  int
	Play  int
	Bonus int
}

func (p *Payout) String() string {
	return fmt.Sprintf("Ante: %d, Play: %d, Bonus: %d\n", p.Ante, p.Play, p.Bonus)
}
