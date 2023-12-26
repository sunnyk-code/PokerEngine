package pokerLogic

import (
	"testing"
)

func TestCardString(t *testing.T) {
	testCases := []struct {
		name     string
		card     Card
		expected string
	}{
		{"Ace of Spades", Card{Rank: Ace, Suit: Spades}, "A♤"},
		{"Two of Hearts", Card{Rank: Two, Suit: Hearts}, "2♡"},
		{"Ten of Diamonds", Card{Rank: Ten, Suit: Diamonds}, "10♦"},
		{"King of Clubs", Card{Rank: King, Suit: Clubs}, "K♣"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.card.String() != tc.expected {
				t.Errorf("String() = %v, want %v", tc.card.String(), tc.expected)
			}
		})
	}
}

func TestCardEquals(t *testing.T) {
	testCases := []struct {
		name     string
		card1    Card
		card2    Card
		expected bool
	}{
		{"Two of Hearts equals Two of Hearts", Card{Rank: Two, Suit: Hearts}, Card{Rank: Two, Suit: Hearts}, true},
		{"Two of Hearts not equals Three of Diamonds", Card{Rank: Two, Suit: Hearts}, Card{Rank: Three, Suit: Diamonds}, false},
		{"Ace of Spades equals Ace of Spades", Card{Rank: Ace, Suit: Spades}, Card{Rank: Ace, Suit: Spades}, true},
		{"Ace of Spades not equals King of Clubs", Card{Rank: Ace, Suit: Spades}, Card{Rank: King, Suit: Clubs}, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.card1.Equals(tc.card2) != tc.expected {
				t.Errorf("Expected card1.Equals(card2) to be %v", tc.expected)
			}
		})
	}
}

func TestGetCardCode(t *testing.T) {
	/* Card Code of A of Spades
	Suit Value: Spades=1. Spades = 1 - 1 = 0.
	Rank Value: Ace=14. In binary, 14 is 1110. << 2, we get 111000, which is 56 in decimal.
	Combining: Using bitwise OR (|) does not change the value. So, the final CardCode for the Ace of Spades is 56.
	*/
	testCases := []struct {
		name     string
		card     Card
		expected CardCode
	}{
		{"Ace of Spades", Card{Rank: Ace, Suit: Spades}, CardCode(56)},     //111000
		{"Two of Hearts", Card{Rank: Two, Suit: Hearts}, CardCode(9)},      //001001
		{"Ten of Diamonds", Card{Rank: Ten, Suit: Diamonds}, CardCode(42)}, //101010
		{"King of Clubs", Card{Rank: King, Suit: Clubs}, CardCode(55)},     //110111
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if cardCode := tc.card.GetCardCode(); cardCode != tc.expected {
				t.Errorf("GetCardCode() = %v, want %v", cardCode, tc.expected)
			}
		})
	}
}

func TestGetAbsoluteValue(t *testing.T) {
	testCases := []struct {
		name     string
		card     Card
		expected int
	}{
		{"Ace of Spades", Card{Rank: Ace, Suit: Spades}, 12},
		{"Two of Hearts", Card{Rank: Two, Suit: Hearts}, 13},
		{"Ten of Diamonds", Card{Rank: Ten, Suit: Diamonds}, 34},
		{"King of Clubs", Card{Rank: King, Suit: Clubs}, 50},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.card.GetAbsoluteValue() != tc.expected {
				t.Errorf("GetAbsoluteValue() = %v, want %v", tc.card.GetAbsoluteValue(), tc.expected)
			}
		})
	}
}
