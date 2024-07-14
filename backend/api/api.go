package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	p "backend/pokerLogic"
)

func helloWorld(c echo.Context) error {

	deck := p.NewDeck()
	myHand := []p.Card{*deck.BorrowRandom(), *deck.BorrowRandom()}
	fmt.Println("myHand: ", myHand)
	communityCards := make([]p.Card, 5)
	communityCards[0] = *deck.BorrowRandom()
	communityCards[1] = *deck.BorrowRandom()
	fmt.Println("communityCards: ", communityCards)

	winPercentage := p.MonteCarloSimulation(2, 10000, p.Hand(myHand), communityCards, deck)

	fmt.Println("Win Percentage: ", winPercentage)

	return c.String(http.StatusOK, "Hello, World!")
}

func ambiguousHands(c echo.Context) error {
	firstHands := [][]p.Card{
		{p.NewCard("A_♤"), p.NewCard("K_♡"), p.NewCard("Q_♣"), p.NewCard("J_♦"), p.NewCard("9_♣")},  // High Card Hand
		{p.NewCard("2_♣"), p.NewCard("4_♣"), p.NewCard("6_♣"), p.NewCard("8_♣"), p.NewCard("10_♣")}, // Flush Hand
		{p.NewCard("5_♦"), p.NewCard("6_♣"), p.NewCard("7_♤"), p.NewCard("8_♡"), p.NewCard("9_♦")},  // Straight Hand
		{p.NewCard("A_♣"), p.NewCard("A_♦"), p.NewCard("K_♣"), p.NewCard("K_♡"), p.NewCard("Q_♤")},  // Two Pair Hand
	}

	secondHands := [][]p.Card{
		{p.NewCard("3_♣"), p.NewCard("3_♦"), p.NewCard("2_♤"), p.NewCard("4_♡"), p.NewCard("5_♤")}, // Pair Hand
		{p.NewCard("3_♡"), p.NewCard("3_♤"), p.NewCard("3_♦"), p.NewCard("2_♣"), p.NewCard("2_♤")}, // Full House Hand
		{p.NewCard("2_♣"), p.NewCard("2_♦"), p.NewCard("2_♤"), p.NewCard("A_♡"), p.NewCard("K_♣")}, // Three of a Kind Hand
		{p.NewCard("2_♣"), p.NewCard("2_♦"), p.NewCard("2_♡"), p.NewCard("2_♤"), p.NewCard("3_♣")}, // Four of a Kind Hand
	}
	for i, _ := range firstHands {
		hand1 := p.Hand(firstHands[i])
		hand2 := p.Hand(secondHands[i])

		handRank1, handValue1 := p.EvaluateFiveCardHand(hand1)
		handRank2, handValue2 := p.EvaluateFiveCardHand(hand2)
		fmt.Printf("%v: %v equaling %v\n", hand1, handRank1, handValue1)
		fmt.Printf("%v: %v equaling %v\n", hand2, handRank2, handValue2)
		if handValue1 > handValue2 {
			fmt.Printf("Hand 1 is stronger!\n\n")
		} else {
			fmt.Printf("Hand 2 is stronger!\n\n")
		}
	}

	return c.String(http.StatusOK, "Hello, World!")
}

func StartServer() {
	e := echo.New()

	e.GET("/", helloWorld)

	e.Logger.Fatal(e.Start(":8080"))
}
