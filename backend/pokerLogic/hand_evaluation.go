package pokerLogic

import (
	"sort"
)

// Hand represents a set of cards
type Hand []Card

// HandRank defines the rank of a hand (e.g., pair, straight)
type HandRank string

const (
	HighCard HandRank      = "High Card"
	Pair                   = "Pair"
	TwoPair                = "Two Pair"
	Trips                  = "Trips"
	Straight               = "Straight"
	Flush                  = "Flush"
	FullHouse              = "Full House"
	Quads            	   = "Quads"
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

	for _, card := range hand {
		rankCounts[byte(card.Rank)]++
		suitCounts[byte(card.Suit)]++
	}

	isFlush := len(suitCounts) == 1
	isStraight := isStraight(hand)

	if isFlush && isStraight {
		if rankCounts[10] == 1 && rankCounts[14] == 1 {
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
		return Quads
	case 3:
		if secondMaxCount >= 2 {
			return FullHouse
		}
		return Trips
	case 2:
		if secondMaxCount == 2 {
			return TwoPair
		}
		return Pair
	}

	return HighCard
}