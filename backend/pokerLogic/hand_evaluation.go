package pokerLogic

import (
	"sort"
	"runtime"
	"sync"
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

// Arguments:
// MonteCarloSimulation performs a Monte Carlo simulation to estimate the probability of winning.
// It takes the number of players, number of iterations, player's hand, community cards, and the deck as arguments.
// numPlayers specifies the number of players in the game (minimum value: 2).
// numIterations specifies the number of iterations to run the simulation (minimum value: 1).
// myHand represents the player's hand and should be a full hand (a slice of 2 Cards).
// communityCards represents the community cards and can have anywhere from 0 to 5 cards filled in the list.
// deck is a pointer to a deck with the appropriate cards removed (ones that are already dealt).
// It returns the estimated probability of winning as a float64.
func MonteCarloSimulation(numPlayers int, numIterations int, myHand Hand, communityCards []Card, deck *Deck) float64 {
	var wins int
	var mu sync.Mutex
	var wg sync.WaitGroup

	numGoroutines := runtime.NumCPU()
	iterationsPerGoroutine := numIterations / numGoroutines

	for g := 0; g < numGoroutines; g++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			localWins := 0
			for i := 0; i < iterationsPerGoroutine; i++ {
				// Reset the deck
				tempDeck := deck.Copy()
				if Simulation(numPlayers, myHand, communityCards, tempDeck) {
					localWins++
				}
			}
			mu.Lock()
			wins += localWins
			mu.Unlock()
		}()
	}

	wg.Wait()
	return float64(wins) / float64(numIterations)
}

func Simulation(numPlayers int, myHand Hand, communityCards []Card, deck *Deck) bool {

	// Deal the rest of the community cards
	numExtraCommunityCards := 5 - len(communityCards)
	extraCommunityCards := make([]Card, numExtraCommunityCards)
	for i := 0; i < numExtraCommunityCards; i++ {
		extraCommunityCards[i] = *deck.BorrowRandom()
	}
	tempCommunityCards := append(communityCards, extraCommunityCards...)

	// Evaluate our hand
	myHand = append(myHand, tempCommunityCards...)
	_, myScore := EvaluateSevenCardHand(myHand)

	for i := 0; i < numPlayers-1; i++ {
		// Deal each player 2 cards & create 7 card hand
		hand := append(tempCommunityCards, *deck.BorrowRandom(), *deck.BorrowRandom())
		// Evaluate the 7 card hand, compare with our hand
		_, score := EvaluateSevenCardHand(hand)
		if score > myScore {
			return false
		}
	}
	return true

}

// Uses the evaluate five card hand function on every possible five card combination of the seven card hand and returns the evaluation for the best possible hand
func EvaluateSevenCardHand(sevenCardHand Hand) (string, int) {

	var bestHandType string
	var highestScore int

	for _, fiveCardHand := range allFiveCardCombinations(sevenCardHand) {
		handType, handValue := EvaluateFiveCardHand(fiveCardHand)
		if handValue > highestScore {
			bestHandType = handType
			highestScore = handValue
		}
	}

	return bestHandType, highestScore

}

func generateCombinations(hand []Card, start, end, index int, combination []Card, result *[][]Card) {
	if index == 5 {
		newCombination := make([]Card, 5)
		copy(newCombination, combination)
		*result = append(*result, newCombination)
		return
	}

	for i := start; i <= end && end-i+1 >= 5-index; i++ {
		combination[index] = hand[i]
		generateCombinations(hand, i+1, end, index+1, combination, result)
	}
}

func allFiveCardCombinations(hand []Card) [][]Card {
	result := make([][]Card, 0)
	combination := make([]Card, 5)
	generateCombinations(hand, 0, len(hand)-1, 0, combination, &result)
	return result
}
