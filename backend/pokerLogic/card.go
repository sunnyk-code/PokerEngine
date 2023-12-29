/*
This file contains the definition of a Card in our application (Based off of pokerlib https://pkg.go.dev/github.com/rbaderts/pokerlib)

We represent cards as 6 bits, with 4 bits for the Rank followed by 2 bits for the Suit.

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
	"strings"
)

type Rank byte

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

type Suit byte

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
		return "Unknown Rank"
	}
}

func (s Suit) String() string {
	switch s {
	case Spades:
		return "♤"
	case Hearts:
		return "♡"
	case Diamonds:
		return "♦"
	case Clubs:
		return "♣"
	default:
		return "Unknown Suit"
	}
}

type Card struct {
	Rank Rank `json:"rank"`
	Suit Suit `json:"suit"`
}

// NewCard creates a new card given its rank and suit
func extractCard(cardStr string) (Rank, Suit) {
	split_card := strings.Split(cardStr, "_")

	raw_rank := strings.ToUpper(string(split_card[0]))
	raw_suit := strings.ToUpper(string(split_card[1]))

	var rank_map = map[string]Rank{
		"2":  Two,
		"3":  Three,
		"4":  Four,
		"5":  Five,
		"6":  Six,
		"7":  Seven,
		"8":  Eight,
		"9":  Nine,
		"10": Ten,
		"J":  Jack,
		"Q":  Queen,
		"K":  King,
		"A":  Ace,
	}

	var suit_map = map[string]Suit{
		"♤": Spades,
		"♦": Diamonds,
		"♡": Hearts,
		"♣": Clubs,
	}

	return rank_map[raw_rank], suit_map[raw_suit]
}

func NewCard(cardStr string) Card {
	rank, suit := extractCard(cardStr)

	return Card{
		Rank: rank,
		Suit: suit,
	}
}

func (c Card) GetAbsoluteValue() int {
	return 13*(int(c.Suit)-1) + int(c.Rank) - 2
}

func (c Card) String() string {
	return fmt.Sprintf("%s%s", c.Rank.String(), c.Suit.String())
}

func (c Card) Equals(other Card) bool {
	return c.Rank == other.Rank && c.Suit == other.Suit
}

type CardCode int

func (c Card) GetCardCode() CardCode {
	rankVal := int(c.Rank) << 2
	suitVal := int(c.Suit - 1)
	return CardCode(rankVal | suitVal)
}

func (c Card) GetRankValue() CardCode {
	return CardCode(c.Rank)
}
