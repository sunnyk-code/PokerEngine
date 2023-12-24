package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	p "pokerLogic"
)

func helloWorld(c echo.Context) error {

	hand := p.Hand{p.NewCard("1h"), p.NewCard("1d"), p.NewCard("3s"), p.NewCard("4h"), p.NewCard("5h")}
	p.SortHand(hand)

	handRank := p.EvaluateHand(hand)

	fmt.Printf("Hand: %v, Rank: %v\n", hand, handRank)

	return c.String(http.StatusOK, "Hello, World!")
}

func StartServer() {
	e := echo.New()

	e.GET("/", helloWorld)

	e.Start(":8080")
}
