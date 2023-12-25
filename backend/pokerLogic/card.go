/*
	This file contains the definition of a Card in our application (Based off of pokerlib)

We represent cards as 6 bits, with 4 bits for the Rank followed by 2 bits for the Rank.

The Suit is represented as follows:

	1: Spades
	2: Hearts
	3: Diamonds
	4: Clubs

The Rank is represented as follows:

	2: Rank 2
	...
	9: Rank 9
	...
	14: Ace Card

Our range for cards will be 0-52: ((Suit-1) * 13) + (Rank-2)
*/
package pokerLogic

import (
	"fmt"
	"strconv"
)

type Rank int

const (
	Two Rank = iota + 2
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

type Suit int

const (
	Spades Suit = iota + 1
	Hearts
	Diamonds
	Clubs
)

func (r Rank) String() string {
	switch r {
	case Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten:
		return strconv.Itoa(int(r))
	case Jack:
		return "J"
	case Queen:
		return "Q"
	case King:
		return "K"
	case Ace:
		return "A"
	default:
		return "Unknown"
	}
}

func (s Suit) String() string {
	switch s {
	case Spades:
		return "Spades"
	case Hearts:
		return "Hearts"
	case Diamonds:
		return "Diamonds"
	case Clubs:
		return "Clubs"
	default:
		return "Unknown Suit"
	}
}

type Card struct {
	Rank Rank `json:"rank"`
	Suit Suit `json:"suit"`
}

func (c Card) GetAbsoluteValue() int {
	return 13*(int(c.Suit)-1) + int(c.Rank) - 2
}

func (c Card) String() string {
	return fmt.Sprintf("%s of %s", c.Rank.String(), c.Suit.String())
}

func (c Card) Equals(other Card) bool {
	return c.Rank == other.Rank && c.Suit == other.Suit
}

type CardCode int

func (c Card) GetCardCode() CardCode {
	rankVal := int(c.Rank) << 4 // Shift the rank over by 4 bits
	suitVal := int(c.Suit - 1)
	return CardCode(rankVal | suitVal)
}

func (c Card) GetRankValue() CardCode {
	return CardCode(c.Rank)
}
