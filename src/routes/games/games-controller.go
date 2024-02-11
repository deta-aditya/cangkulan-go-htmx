package games

import (
	"context"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"example.com/htmx-example/src/services/cards"
	"example.com/htmx-example/src/services/games"
	"github.com/go-chi/chi/v5"
)

type contextKey int

const (
	gameKey contextKey = iota
)

type Dependencies struct {
	gameRepository games.GameRepository
}

type gameState struct {
	Game    games.Game
	Players map[int]player
}

type player struct {
	Name  string
	Hands []cardState
}

type cardState struct {
	RankFmt    string
	SuitFmt    template.HTML
	ColorClass string
	PlayUri    string
}

var rankFmtMap = map[cards.Rank]string{
	cards.RankTwo:   "2",
	cards.RankThree: "3",
	cards.RankFour:  "4",
	cards.RankFive:  "5",
	cards.RankSix:   "6",
	cards.RankSeven: "7",
	cards.RankEight: "8",
	cards.RankNine:  "9",
	cards.RankTen:   "10",
	cards.RankJack:  "J",
	cards.RankQueen: "Q",
	cards.RankKing:  "K",
	cards.RankAce:   "A",
}

var suitFmtMap = map[cards.Suit]template.HTML{
	cards.SuitDiamond: "&diams;",
	cards.SuitClub:    "&clubs;",
	cards.SuitHeart:   "&hearts;",
	cards.SuitSpade:   "&spades;",
}

func gamesController(dependencies Dependencies) func(chi.Router) {
	gameContext := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			gameIDRaw := chi.URLParam(r, "gameID")
			gameID, err := strconv.Atoi(gameIDRaw)
			if err != nil {
				//
			}

			game, err := dependencies.gameRepository.GetByID(int32(gameID))
			if err != nil {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}

			ctx := context.WithValue(r.Context(), gameKey, game)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}

	getBaseGame := func(w http.ResponseWriter, r *http.Request) {
		game, err := getGameFromContext(r)
		if err != nil {

		}

		baseTemplate := template.Must(template.ParseFiles(
			"src/templates/game.html",
			"src/templates/player.html",
			"src/templates/card.html",
		))
		state := initializeState(game)

		err = baseTemplate.Execute(w, state)
		if err != nil {
			log.Println(err.Error())
		}
	}

	return func(router chi.Router) {
		router.Use(gameContext)
		router.Get("/", getBaseGame)
	}
}

func getGameFromContext(r *http.Request) (games.Game, error) {
	ctx := r.Context()
	game, ok := ctx.Value(gameKey).(games.Game)
	if !ok {
		return game, errors.New("cannot get game from context")
	}
	return game, nil
}

func initializeState(g games.Game) gameState {
	s := gameState{
		Game:    g,
		Players: make(map[int]player),
	}

	for k, h := range g.Hands {
		cs := make([]cardState, 0, len(h))

		for _, c := range h {
			colorClass := "mono"
			if c.Suit == cards.SuitHeart || c.Suit == cards.SuitDiamond {
				colorClass = "color"
			}

			cs = append(cs, cardState{
				RankFmt:    rankFmtMap[c.Rank],
				SuitFmt:    suitFmtMap[c.Suit],
				ColorClass: colorClass,
				PlayUri:    fmt.Sprintf("/games/%d/cards/%d-%d", *g.ID, c.Rank, c.Suit),
			})
		}

		s.Players[k] = player{
			Name:  fmt.Sprintf("Player %v", k),
			Hands: cs,
		}
	}

	return s
}

func Register(router *chi.Mux, gameRepository games.GameRepository) {
	router.Route("/games/{gameID}", gamesController(Dependencies{
		gameRepository: gameRepository,
	}))
}
