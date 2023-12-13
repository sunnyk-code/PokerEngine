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

func isFlush(hand Hand) bool {
	suit := hand[0].Suit
	for _, card := range hand[1:] {
		if card.Suit != suit {
			return false
		}
	}
	return true
}

func isStraight(hand Hand) bool {
	var ranks []int
	for _, card := range hand {
		ranks = append(ranks, int(card.Rank))
	}
	sort.Ints(ranks)

	for i := 1; i < len(ranks); i++ {
		if ranks[i] != ranks[i-1]+1 {
			return false
		}
	}
	return true
}

func countRanks(hand Hand) map[byte]int {
	rankCounts := make(map[byte]int)
	for _, card := range hand {
		rankCounts[card.Rank]++
	}
	return rankCounts
}

func evaluateByCount(rankCounts map[byte]int) HandRank {
	pairs := 0
	threeOfAKindFound := false

	for _, count := range rankCounts {
		switch count {
		case 4:
			return FourOfAKind
		case 3:
			threeOfAKindFound = true
		case 2:
			pairs++
		}
	}

	if threeOfAKindFound && pairs > 0 {
		return FullHouse
	}
	if threeOfAKindFound {
		return ThreeOfAKind
	}
	if pairs == 2 {
		return TwoPair
	}
	if pairs == 1 {
		return Pair
	}

	return HighCard
}

func EvaluateHand(hand Hand) HandRank {
	rankCounts := countRanks(hand)
	if isFlush(hand) && isStraight(hand) {
		if hand[0].Rank == 'a' {
			return RoyalFlush
		}
		return StraightFlush
	}

	if isFlush(hand) {
		return Flush
	}
	if isStraight(hand) {
		return Straight
	}

	return evaluateByCount(rankCounts)
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
	hand := Hand{NewCard("1h"), NewCard("2h"), NewCard("3h"), NewCard("4h"), NewCard("5h")}
	SortHand(hand)

	// Evaluate the hand
	handRank := EvaluateHand(hand)

	fmt.Printf("Hand: %v, Rank: %v\n", hand, handRank)
}
