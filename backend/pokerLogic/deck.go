package pokerLogic

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgryski/go-pcgr"
)

type Deck struct {
	cardArray [52]*Card
	cardIndex map[Card]int
	size      int
	rnd       *pcgr.Rand
}

func (deck *Deck) appendToBottom(card Card) {
	deck.cardArray[deck.size] = &card
	deck.cardIndex[card] = deck.size
	deck.size += 1
}


func (d *Deck) removeCard(card Card) {
	index, exists := d.cardIndex[card]
	if !exists {
		return
	}
	d.cardArray[index] = nil
	d.size -= 1
}

func NewDeck() *Deck {
	deck := new(Deck)
	deck.rnd = &pcgr.Rand{uint64(time.Now().UnixNano()), 0x00004443}
	deck.cardIndex = make(map[Card]int, 0)

	index := 0
	for rank := Two; rank <= Ace; rank++ {
		for suit := 1; suit <= 4; suit++ {
			card := Card{Rank(rank), Suit(suit)}
			deck.appendToBottom(card)
			index++
		}
	}
	return deck
}

func (deck *Deck) BorrowRandom() *Card {
	// var card *Card
	var randomCard *Card
	for {
		index := int(deck.rnd.Bound(uint32(52)))
		randomCard = deck.cardArray[index]
		if randomCard == nil {
			continue
		}
		deck.removeCard(*randomCard)
		break
	}

	return randomCard
}

func (deck *Deck) ReturnCard(card Card) error {
	index := deck.cardIndex[card]
	if deck.cardArray[index] != nil {
		return errors.New(fmt.Sprintf("This card - %v, was not taken from the deck\n", card))
	}
	deck.cardArray[index] = &card
	return nil
}

func (d *Deck) Copy() *Deck {
	newDeck := new(Deck)
	newDeck.rnd = &pcgr.Rand{uint64(time.Now().UnixNano()), 0x00004443}
	newDeck.cardIndex = make(map[Card]int, 0)
	newDeck.size = d.size

	for i, card := range d.cardArray {
		if card != nil {
			newCard := *card
			newDeck.cardArray[i] = &newCard
			newDeck.cardIndex[newCard] = i
		}
	}

	return newDeck
}
