package models

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	tests := []struct {
		name     string
		input    Hand
		expected Hand
	}{
		{
			name: "simple hand",
			input: Hand{cards: []Card{
				{value: 14, rank: "A", suit: "D"},
				{value: 2, rank: "2", suit: "H"},
				{value: 7, rank: "7", suit: "C"},
			}},
			expected: Hand{cards: []Card{
				{value: 2, rank: "2", suit: "H"},
				{value: 7, rank: "7", suit: "C"},
				{value: 14, rank: "A", suit: "D"},
			}},
		},
		{
			name: "with pair",
			input: Hand{cards: []Card{
				{value: 2, rank: "2", suit: "H"},
				{value: 7, rank: "7", suit: "C"},
				{value: 2, rank: "2", suit: "D"},
			}},
			expected: Hand{cards: []Card{
				{value: 2, rank: "2", suit: "D"},
				{value: 2, rank: "2", suit: "H"},
				{value: 7, rank: "7", suit: "C"},
			}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.input.sort()

			if !reflect.DeepEqual(test.input, test.expected) {
				t.Errorf("FAIL: expected %v output: %v", test.expected, test.input)
			}
		})
	}

}

func TestIsConsecutive(t *testing.T) {
	tests := []struct {
		name     string
		input    Hand
		expected bool
	}{
		{
			name: "consecutive with low ace",
			input: Hand{cards: []Card{
				{value: 14, rank: "A", suit: "D"},
				{value: 2, rank: "2", suit: "H"},
				{value: 3, rank: "3", suit: "C"},
			}},
			expected: true,
		},
		{
			name: "consecutive with high ace",
			input: Hand{cards: []Card{
				{value: 14, rank: "A", suit: "D"},
				{value: 13, rank: "K", suit: "H"},
				{value: 12, rank: "Q", suit: "C"},
			}},
			expected: true,
		},
		{
			name: "simple non-consecutive",
			input: Hand{cards: []Card{
				{value: 8, rank: "8", suit: "D"},
				{value: 2, rank: "2", suit: "H"},
				{value: 6, rank: "6", suit: "C"},
			}},
			expected: false,
		},
		{
			name: "simple consecutive",
			input: Hand{cards: []Card{
				{value: 8, rank: "8", suit: "D"},
				{value: 7, rank: "7", suit: "H"},
				{value: 6, rank: "6", suit: "C"},
			}},
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := test.input.isConsecutive()

			if output != test.expected {
				t.Errorf("FAIL: expected %t output %t", test.expected, output)
			}

		})
	}
}

func TestIsOnePair(t *testing.T) {
	tests := []struct {
		name     string
		input    Hand
		expected bool
	}{
		{
			name: "positive test",
			input: Hand{cards: []Card{
				{value: 10, rank: "10", suit: "C"},
				{value: 10, rank: "10", suit: "D"},
				{value: 2, rank: "2", suit: "H"}},
			},
			expected: true,
		},
		{
			name: "negative test",
			input: Hand{cards: []Card{
				{value: 10, rank: "10", suit: "C"},
				{value: 6, rank: "6", suit: "D"},
				{value: 2, rank: "2", suit: "H"}},
			},
			expected: false,
		},
		{
			name: "three of a kind",
			input: Hand{cards: []Card{
				{value: 10, rank: "10", suit: "C"},
				{value: 10, rank: "10", suit: "D"},
				{value: 10, rank: "10", suit: "H"}},
			},
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := test.input.isOnePair()
			if output != test.expected {
				t.Errorf("FAIL: expected %t output %t", test.expected, output)
			}
		})
	}

}

func TestIsFlush(t *testing.T) {
	tests := []struct {
		name     string
		input    Hand
		expected bool
	}{
		{
			name: "positive test",
			input: Hand{cards: []Card{
				{value: 10, rank: "10", suit: "C"},
				{value: 7, rank: "7", suit: "C"},
				{value: 2, rank: "2", suit: "C"}},
			},
			expected: true,
		},
		{
			name: "negative test",
			input: Hand{cards: []Card{
				{value: 10, rank: "10", suit: "C"},
				{value: 6, rank: "6", suit: "D"},
				{value: 2, rank: "2", suit: "H"}},
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := test.input.isFlush()
			if output != test.expected {
				t.Errorf("FAIL: expected %t output %t", test.expected, output)
			}
		})
	}

}

func TestIsStraight(t *testing.T) {
	tests := []struct {
		name     string
		input    Hand
		expected bool
	}{
		{
			name: "positive test",
			input: Hand{cards: []Card{
				{value: 10, rank: "10", suit: "C"},
				{value: 9, rank: "9", suit: "H"},
				{value: 8, rank: "8", suit: "D"}},
			},
			expected: true,
		},
		{
			name: "negative test",
			input: Hand{cards: []Card{
				{value: 10, rank: "10", suit: "C"},
				{value: 6, rank: "6", suit: "D"},
				{value: 2, rank: "2", suit: "H"}},
			},
			expected: false,
		},
		{
			name: "postive test with high ace",
			input: Hand{cards: []Card{
				{value: 14, rank: "A", suit: "C"},
				{value: 13, rank: "K", suit: "D"},
				{value: 12, rank: "Q", suit: "H"}},
			},
			expected: true,
		},
		{
			name: "postive test with low ace",
			input: Hand{cards: []Card{
				{value: 14, rank: "A", suit: "C"},
				{value: 2, rank: "2", suit: "D"},
				{value: 3, rank: "3", suit: "H"}},
			},
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := test.input.isStraight()
			if output != test.expected {
				t.Errorf("FAIL: expected %t output %t", test.expected, output)
			}
		})
	}

}

func TestIsThreeOfAKind(t *testing.T) {
	tests := []struct {
		name     string
		input    Hand
		expected bool
	}{
		{
			name: "positive test",
			input: Hand{cards: []Card{
				{value: 10, rank: "10", suit: "C"},
				{value: 10, rank: "10", suit: "H"},
				{value: 10, rank: "10", suit: "D"}},
			},
			expected: true,
		},
		{
			name: "negative test",
			input: Hand{cards: []Card{
				{value: 10, rank: "10", suit: "C"},
				{value: 6, rank: "6", suit: "D"},
				{value: 2, rank: "2", suit: "H"}},
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := test.input.isThreeOfAKind()
			if output != test.expected {
				t.Errorf("FAIL: expected %t output %t", test.expected, output)
			}
		})
	}

}

func TestIsStraightFlush(t *testing.T) {
	tests := []struct {
		name     string
		input    Hand
		expected bool
	}{
		{
			name: "positive test",
			input: Hand{cards: []Card{
				{value: 10, rank: "10", suit: "S"},
				{value: 9, rank: "9", suit: "S"},
				{value: 8, rank: "8", suit: "S"}},
			},
			expected: true,
		},
		{
			name: "negative test",
			input: Hand{cards: []Card{
				{value: 10, rank: "10", suit: "C"},
				{value: 6, rank: "6", suit: "D"},
				{value: 2, rank: "2", suit: "H"}},
			},
			expected: false,
		},
		{
			name: "postive test with high ace",
			input: Hand{cards: []Card{
				{value: 14, rank: "A", suit: "S"},
				{value: 13, rank: "K", suit: "S"},
				{value: 12, rank: "Q", suit: "S"}},
			},
			expected: true,
		},
		{
			name: "postive test with low ace",
			input: Hand{cards: []Card{
				{value: 14, rank: "A", suit: "S"},
				{value: 2, rank: "2", suit: "S"},
				{value: 3, rank: "3", suit: "S"}},
			},
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := test.input.isStraightFlush()
			if output != test.expected {
				t.Errorf("FAIL: expected %t output %t", test.expected, output)
			}
		})
	}

}

func TestIsRoyalFlush(t *testing.T) {
	tests := []struct {
		name     string
		input    Hand
		expected bool
	}{
		{
			name: "positive test with spades",
			input: Hand{cards: []Card{
				{value: 12, rank: "Q", suit: "S"},
				{value: 13, rank: "K", suit: "S"},
				{value: 14, rank: "A", suit: "S"}},
			},
			expected: true,
		},
		{
			name: "positive test with diamonds",
			input: Hand{cards: []Card{
				{value: 12, rank: "Q", suit: "D"},
				{value: 13, rank: "K", suit: "D"},
				{value: 14, rank: "A", suit: "D"}},
			},
			expected: true,
		},
		{
			name: "negative test with matching suits",
			input: Hand{cards: []Card{
				{value: 2, rank: "2", suit: "S"},
				{value: 3, rank: "3", suit: "S"},
				{value: 14, rank: "A", suit: "S"}},
			},
			expected: false,
		},
		{
			name: "negative test with matching suits",
			input: Hand{cards: []Card{
				{value: 2, rank: "2", suit: "S"},
				{value: 3, rank: "3", suit: "H"},
				{value: 14, rank: "A", suit: "C"}},
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := test.input.isRoyalFlush()
			if output != test.expected {
				t.Errorf("FAIL: expected %t output %t", test.expected, output)
			}
		})
	}

}
