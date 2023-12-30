package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"math/rand"
	"time"

	p "pokerLogic"
)

func helloWorld(c echo.Context) error {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Define slices for card numbers and suits
	numbers := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	suits := []string{"♣", "♦", "♡", "♤"}

	for i := 0; i < 150; i++ {
		var selectedCards []p.Card

		for j := 0; j < 5; j++ {
			number := numbers[rand.Intn(len(numbers))]
			suit := suits[rand.Intn(len(suits))]
			selectedCards = append(selectedCards, p.NewCard(number+"_"+suit))
		}

		hand := p.Hand(selectedCards)
		handRank, handValue := p.EvaluateHand(hand)
		fmt.Printf("Hand: %v, Rank: %v: %v\n", hand, handRank, handValue)
	}
	return c.String(http.StatusOK, "Hello, World!")
}

func ambiguousHands(c echo.Context) error {
	firstHands := [][]p.Card{
		{p.NewCard("A_♤"), p.NewCard("K_♡"), p.NewCard("Q_♣"), p.NewCard("J_♦"), p.NewCard("9_♣")}, // High Card Hand
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

		handRank1, handValue1 := p.EvaluateHand(hand1)
		handRank2, handValue2 := p.EvaluateHand(hand2)
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

	e.GET("/", ambiguousHands)

	e.Logger.Fatal(e.Start(":8080"))
}
