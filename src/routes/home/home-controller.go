package home

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	cardsPkg "example.com/htmx-example/src/services/cards"
	"example.com/htmx-example/src/services/games"
	"github.com/go-chi/chi/v5"
)

type Dependencies struct {
	gameRepository games.GameRepository
}

func homeController(dependencies Dependencies) func(chi.Router) {
	getHome := func(w http.ResponseWriter, r *http.Request) {
		htmlTemplate := template.Must(template.ParseFiles("src/templates/home.html"))
		htmlTemplate.Execute(w, nil)
	}

	getCreateGameModal := func(w http.ResponseWriter, r *http.Request) {
		template.Must(template.ParseFiles("src/templates/create-game-modal.html")).Execute(w, nil)
	}

	createGame := func(w http.ResponseWriter, r *http.Request) {
		players, err := strconv.Atoi(r.PostFormValue("players"))
		if err != nil {
			// TODO: handle error
		}

		cards, err := strconv.Atoi(r.PostFormValue("cards"))
		if err != nil {
			// TODO: handle error
		}

		newGame := games.NewGame(int32(players), int32(cards))
		newGame.DistributeHand(cardsPkg.GenerateShuffled())

		dependencies.gameRepository.Save(newGame)
		newGameId := dependencies.gameRepository.GetRecentyCreated().ID

		w.Header().Add("HX-Location", fmt.Sprintf("/games/%v", *newGameId))
	}

	return func(router chi.Router) {
		router.Get("/", getHome)
		router.Get("/games", getCreateGameModal)
		router.Post("/games", createGame)
	}
}

func Register(router *chi.Mux, gameRepository games.GameRepository) {
	router.Route("/home", homeController(Dependencies{
		gameRepository: gameRepository,
	}))
}
