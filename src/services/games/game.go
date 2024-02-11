package games

import "example.com/htmx-example/src/services/cards"

type Game struct {
	ID      *int32
	Players int32
	Cards   int32
	Hands   map[int][]cards.Card
	Deck    []cards.Card
}

func NewGame(players int32, cardsNum int32) Game {
	return Game{
		ID:      nil,
		Players: players,
		Cards:   cardsNum,
		Hands:   make(map[int][]cards.Card),
		Deck:    make([]cards.Card, 0),
	}
}

func (g *Game) SetId(id int32) {
	g.ID = &id
}

func (g *Game) DistributeHand(deck []cards.Card) {
	for i, c := range deck {
		handedCardsMax := g.Players * g.Cards
		playerIndex := i%int(g.Players) + 1

		if i < int(handedCardsMax) {
			g.Hands[playerIndex] = append(g.Hands[playerIndex], c)
		} else {
			g.Deck = append(g.Deck, c)
		}
	}
}
