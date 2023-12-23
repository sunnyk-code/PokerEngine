package main

import (
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
type HandRank string

const (
	HighCard      HandRank = "High Card"
	Pair                   = "Pair"
	TwoPair                = "Two Pair"
	ThreeOfAKind           = "Three of a Kind"
	Straight               = "Straight"
	Flush                  = "Flush"
	FullHouse              = "Full House"
	FourOfAKind            = "Four of a Kind"
	StraightFlush          = "Straight Flush"
	RoyalFlush             = "Royal Flush"
)

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

func EvaluateHand(hand Hand) HandRank {
	rankCounts := make(map[byte]int)
	suitCounts := make(map[byte]int)
	var ranks []int

	for _, card := range hand {
		rankCounts[card.Rank]++
		suitCounts[card.Suit]++
		ranks = append(ranks, int(card.Rank))
	}

	sort.Ints(ranks)

	isFlush := len(suitCounts) == 1

	isStraight := isStraight(hand)

	if isFlush && isStraight {
		if ranks[0] == int('T') {
			return RoyalFlush
		}
		return StraightFlush
	}
	if isFlush {
		return Flush
	}
	if isStraight {
		return Straight
	}

	var maxCount, secondMaxCount int
	for _, count := range rankCounts {
		if count > maxCount {
			secondMaxCount = maxCount
			maxCount = count
		} else if count > secondMaxCount {
			secondMaxCount = count
		}
	}

	switch maxCount {
	case 4:
		return FourOfAKind
	case 3:
		if secondMaxCount >= 2 {
			return FullHouse
		}
		return ThreeOfAKind
	case 2:
		if secondMaxCount == 2 {
			return TwoPair
		}
		return Pair
	}

	return HighCard
}

func SortHand(hand Hand) {
	sort.Slice(hand, func(i, j int) bool {
		if hand[i].Rank == hand[j].Rank {
			return hand[i].Suit < hand[j].Suit
		}
		return hand[i].Rank < hand[j].Rank
	})
}

// func main() {
// 	hand := Hand{NewCard("1h"), NewCard("1d"), NewCard("3s"), NewCard("4h"), NewCard("5h")}
// 	SortHand(hand)

// 	handRank := EvaluateHand(hand)

// 	fmt.Printf("Hand: %v, Rank: %v\n", hand, handRank)
// }
