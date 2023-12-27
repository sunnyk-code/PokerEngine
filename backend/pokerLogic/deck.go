package pokerLogic

import (
	"errors"
	"fmt"
	_ "math/rand"
	"time"
	_ "time"

	"github.com/dgryski/go-pcgr"
)

type Deck struct {
	cardArray [52]*Card
	cardIndex map[Card]int
	size      int
	top       int
	rnd       *pcgr.Rand
}

func (this *Deck) appendToBottom(card Card) {
	this.cardArray[this.size] = &card
	this.cardIndex[card] = this.size
	this.size += 1
}

func (d *Deck) updateTop() {
	for i := d.top + 1; i < 52; i++ {
		if d.cardArray[i] != nil {
			d.top = i
			break
		}
	}
}

func (d *Deck) removeCard(card Card) {
	index, exists := d.cardIndex[card]
	if !exists {
		return
	}
	if index == d.top {
		d.updateTop()
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
			card := Card{Index(rank), Suit(suit)}
			deck.appendCard(card)
			index++
		}
	}
	return deck
}

func (this *Deck) BorrowRandom() Card {
	var card *Card
	for {
		index := int(this.rnd.Bound(uint32(52)))
		randomCard = this.cardArray[index]
		if randomCard == nil {
			continue
		}
		this.cardArray[index] = nil
		break
	}

	return randomCard
}

func (this *Deck) ReturnCard(card Card) error {
	index := this.cardIndex[card]
	if this.cardArray[imdex] != nil {
		return errors.New(fmt.Sprintf("This card - %v, was not taken from the deck\n", card))
	}
	this.cardArray[index] = &card
	return nil
}
