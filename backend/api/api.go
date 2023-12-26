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

	for i := 0; i < 10; i++ {
		var selectedCards []p.Card

		for j := 0; j < 5; j++ {
			number := numbers[rand.Intn(len(numbers))]
			suit := suits[rand.Intn(len(suits))]
			selectedCards = append(selectedCards, p.NewCard(number+"_"+suit))
		}

		hand := p.Hand(selectedCards)
		handRank := p.EvaluateHand(hand)
		fmt.Printf("Hand: %v, Rank: %v\n", hand, handRank)
	}
	return c.String(http.StatusOK, "Hello, World!")
}

func StartServer() {
	e := echo.New()

	e.GET("/", helloWorld)

	e.Logger.Fatal(e.Start(":8080"))
}
