package cards

import (
	"math/rand"
	"time"
)

type Suit int

const (
	SuitDiamond Suit = iota
	SuitClub
	SuitHeart
	SuitSpade
)

var Suits = []Suit{SuitDiamond, SuitClub, SuitHeart, SuitSpade}

type Rank int

const (
	RankTwo Rank = iota
	RankThree
	RankFour
	RankFive
	RankSix
	RankSeven
	RankEight
	RankNine
	RankTen
	RankJack
	RankQueen
	RankKing
	RankAce
)

var Ranks = []Rank{
	RankTwo,
	RankThree,
	RankFour,
	RankFive,
	RankSix,
	RankSeven,
	RankEight,
	RankNine,
	RankTen,
	RankJack,
	RankQueen,
	RankKing,
	RankAce,
}

type Card struct {
	Suit Suit
	Rank Rank
}

func GenerateShuffled() []Card {
	cards := make([]Card, 0, 52)

	for s := range Suits {
		for r := range Ranks {
			cards = append(cards, Card{
				Suit: Suit(s),
				Rank: Rank(r),
			})
		}
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := len(cards) - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		cards[i], cards[j] = cards[j], cards[i]
	}

	return cards
}
