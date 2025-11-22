package models

import "fmt"

const RED = "\x1b[31m"
const BLACK = "\x1b[30m"
const RESET = "\x1b[0m"

type Card struct {
	value uint8
	rank  string
	suit  string
}

func (c *Card) String() []string {

	if c.suit == "H" || c.suit == "D" {
		return []string{
			fmt.Sprintf("%s-----%s", RED, RESET),
			fmt.Sprintf("%s│%-2s%s│%s", RED, c.rank, c.suit, RESET),
			fmt.Sprintf("%s-----%s", RED, RESET),
		}

	}

	return []string{
		fmt.Sprintf("%s-----%s", BLACK, RESET),
		fmt.Sprintf("%s│%-2s%s│%s", BLACK, c.rank, c.suit, RESET),
		fmt.Sprintf("%s-----%s", BLACK, RESET),
	}

}
