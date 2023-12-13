package main

import (
	"fmt"
	"sort"
)

// Card represents a playing card with a rank and a suit
type Card struct {
	Rank byte
	Suit byte
}

// Hand represents a set of cards
type Hand []Card

// NewCard creates a new card given its rank and suit
func NewCard(cardStr string) Card {
	return Card{
		Rank: cardStr[0],
		Suit: cardStr[1],
	}
}

// HandRank defines the rank of a hand (e.g., pair, straight)
type HandRank int

const (
	HighCard HandRank = iota
	Pair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
	RoyalFlush
)

// EvaluateHand evaluates the hand and returns its rank
func EvaluateHand(hand Hand) HandRank {
	// This function should be filled with logic to evaluate the hand.
	// This is a placeholder for simplicity.
	return HighCard // Replace with actual hand evaluation logic
}

// SortHand sorts the hand by rank and suit
func SortHand(hand Hand) {
	sort.Slice(hand, func(i, j int) bool {
		if hand[i].Rank == hand[j].Rank {
			return hand[i].Suit < hand[j].Suit
		}
		return hand[i].Rank < hand[j].Rank
	})
}

func main() {
	// Example: Create a hand
	hand := Hand{NewCard("Qh"), NewCard("9h"), NewCard("Th"), NewCard("Jh"), NewCard("Kh")}

	// Sort the hand (optional, but can be helpful for evaluation)
	SortHand(hand)

	// Evaluate the hand
	handRank := EvaluateHand(hand)

	fmt.Printf("Hand: %v, Rank: %v\n", hand, handRank)
}
