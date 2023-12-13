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
	rankCounts := make(map[byte]int)
	for _, card := range hand {
		rankCounts[card.Rank]++
	}

	pairs := 0
	for _, count := range rankCounts {
		switch count {
		case 2:
			pairs++
		case 3:
			return ThreeOfAKind
			// Add cases for other hand types
		}
	}

	if pairs == 1 {
		return Pair
	} else if pairs == 2 {
		return TwoPair
	}

	return HighCard
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
	hand := Hand{NewCard("Qh"), NewCard("9d"), NewCard("1s"), NewCard("1d"), NewCard("Qh")}
	SortHand(hand)

	// Evaluate the hand
	handRank := EvaluateHand(hand)

	fmt.Printf("Hand: %v, Rank: %v\n", hand, handRank)
}
