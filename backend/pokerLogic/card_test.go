package pokerLogic

import (
	"testing"
)

func TestCardString(t *testing.T) {
	card := Card{Rank: Ace, Suit: Spades}
	expected := "A of Spades"
	if card.String() != expected {
		t.Errorf("String() = %v, want %v", card.String(), expected)
	}
}

func TestCardEquals(t *testing.T) {
	card1 := Card{Rank: Two, Suit: Hearts}
	card2 := Card{Rank: Two, Suit: Hearts}
	card3 := Card{Rank: Three, Suit: Diamonds}

	if !card1.Equals(card2) {
		t.Errorf("Expected card1 to be equal to card2")
	}
	if card1.Equals(card3) {
		t.Errorf("Expected card1 not to be equal to card3")
	}
}

func TestGetCardCode(t *testing.T) {
	card := Card{Rank: Ace, Suit: Spades}
	/* Card Code of A of Spades
	Suit Value: Spades=1. Spades = 1 - 1 = 0.
	Rank Value: Ace=14. In binary, 14 is 1110. << 4, we get 1110000, which is 224 in decimal.
	Combining: Using bitwise OR (|) does not change the value. So, the final CardCode for the Ace of Spades is 224.
	*/
	expected := CardCode(224) // Example value for A of Spaces

	if card.GetCardCode() != expected {
		t.Errorf("GetCardCode() = %v, want %v", card.GetCardCode(), expected)
	}
}
