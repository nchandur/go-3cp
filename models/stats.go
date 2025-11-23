package models

type Stats struct {
	Wins   map[string]uint64
	Losses map[string]uint64
	Pushes map[string]uint64
}

func NewStats() *Stats {
	wins := make(map[string]uint64)
	losses := make(map[string]uint64)
	pushes := make(map[string]uint64)

	for _, hand := range []string{"High Card", "One Pair", "Flush", "Straight", "Three of a Kind", "Straight Flush", "Royal Flush"} {
		wins[hand] = 0
		losses[hand] = 0
		pushes[hand] = 0
	}

	return &Stats{Wins: wins, Losses: losses, Pushes: pushes}

}
