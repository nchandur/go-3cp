package models

type Wager struct {
	Ante  uint64
	Play  uint64
	Bonus uint64
}

func NewWager(ante, bonus uint64) *Wager {
	return &Wager{
		Ante:  ante,
		Bonus: bonus,
	}
}
