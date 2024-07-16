package api

import (
	"fmt"
	"net/http"


	p "backend/pokerLogic"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {

	deck := p.NewDeck()
	// myHand := []p.Card{*deck.BorrowRandom(), *deck.BorrowRandom()}
	myHand := []p.Card{
		p.NewCard("7_♤"),
		p.NewCard("2_♡"),
	}
	fmt.Println("myHand: ", myHand)
	deck.RemoveCard(p.NewCard("7_♤"))
	deck.RemoveCard(p.NewCard("2_♡"))
	communityCards := make([]p.Card, 5)
	communityCards[0] = *deck.BorrowRandom()
	communityCards[1] = *deck.BorrowRandom()
	fmt.Println("communityCards: ", communityCards)

	winPercentage := p.MonteCarloSimulation(2, 100000, p.Hand(myHand), communityCards, deck)

	fmt.Println("Win Percentage: ", winPercentage)

}


func StartServer() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/winning-percentage", ApiKeyMiddleware(WinningPercentageHandler))

	http.ListenAndServe(":8080", nil)
}
