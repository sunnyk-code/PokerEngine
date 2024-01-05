package pokerLogic

import (
	"sort"
)

// Hand represents a set of cards
type Hand []Card

// HandRank defines the rank of a hand (e.g., pair, straight)
type HandRank int

const (
	HighCard HandRank = iota * 1000000000
	Pair
	TwoPair
	Trips
	Straight
	Flush
	FullHouse
	Quads
	StraightFlush
	RoyalFlush
)

func isStraight(ranks []int) bool {
	for i := 1; i < len(ranks)-1; i++ {
		if ranks[i] != ranks[i-1]+1 {
			return false
		}
	}
	if ranks[len(ranks)-1] != ranks[len(ranks)-2]+1 {
		if ranks[len(ranks)-1] == 14 {
			if ranks[0] != 2 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func EvaluateFiveCardHand(hand Hand) (string, int) {
	var ranks []int
	for _, card := range hand {
		ranks = append(ranks, int(card.Rank))
	}
	sort.Ints(ranks)
	
	rankCounts := make(map[int]int)
	suitCounts := make(map[int]int)

	for _, card := range hand {
		rankCounts[int(card.Rank)]++
		suitCounts[int(card.Suit)]++
	}

	isFlush := len(suitCounts) == 1
	isStraight := isStraight(ranks)

	if isFlush && isStraight {
		if ranks[0] == 10 {
			return "RoyalFlush", int(RoyalFlush)
		} else if ranks[0] == 2 && ranks[len(ranks)-1] == 14 {
			return "StraightFlush", int(StraightFlush) + 5
		}
		return "StraightFlush", int(StraightFlush) + ranks[len(ranks)-1]
	}

	if isFlush {
		return "Flush", int(Flush) + ranks[len(ranks)-1]
	}

	if isStraight {
		if ranks[0] == 2 && ranks[len(ranks)-1] == 14 {
			return "Straight", int(Straight) + 5
		}
		return "Straight", int(Straight) + ranks[len(ranks)-1]
	}

	var maxCountRank, secondMaxCountRank int
	var maxCount, secondMaxCount int
	for rank, count := range rankCounts {
		if count > maxCount {
			secondMaxCount = maxCount
			secondMaxCountRank = maxCountRank
			maxCount = count
			maxCountRank = rank
		} else if count > secondMaxCount {
			secondMaxCount = count
			secondMaxCountRank = rank
		}
	}

	switch maxCount {
	case 4:
		kicker := 0
		for _, rank := range ranks {
			if rank != maxCountRank {
				kicker = rank
				break
			}
		}
		return "Quads", int(Quads) + maxCountRank*10000000 + kicker
	case 3:
		if secondMaxCount == 2 {
			return "FullHouse", int(FullHouse) + maxCountRank*10000000 + secondMaxCountRank
		}
		var kickers []int
		for _, rank := range ranks {
			if rank != maxCountRank {
				kickers = append(kickers, rank)
			}
		}
		sort.Sort(sort.Reverse(sort.IntSlice(kickers)))
		return "Trips", int(Trips) + maxCountRank*10000000 + kickers[0]*100000 + kickers[1]*1000
	case 2:
		if secondMaxCount == 2 {
			kicker := 0
			for _, rank := range ranks {
				if rank != maxCountRank && rank != secondMaxCountRank {
					kicker = rank
					break
				}
			}
			return "TwoPair", int(TwoPair) + max(maxCountRank, secondMaxCountRank)*10000000 + min(maxCountRank, secondMaxCountRank)*100000 + kicker*1000
		}
		var kickers []int
		for _, rank := range ranks {
			if rank != maxCountRank {
				kickers = append(kickers, rank)
			}
		}
		sort.Sort(sort.Reverse(sort.IntSlice(kickers)))
		return "Pair", int(Pair) + maxCountRank*10000000 + kickers[0]*100000 + kickers[1]*1000 + kickers[2]*10
	}

	return "HighCard", int(HighCard) + ranks[4]*10000000 + ranks[3]*100000 + ranks[2]*1000 + ranks[1]*10 + ranks[0]
}