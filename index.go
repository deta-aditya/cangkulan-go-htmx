package main

import (
	"log"
	"net/http"

	gr "example.com/htmx-example/src/routes/games"
	"example.com/htmx-example/src/routes/home"
	"example.com/htmx-example/src/services/games"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	gamesRepo := games.NewInMemoryGameRepository()

	router.Handle("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	router.Handle("/", http.RedirectHandler("/home", http.StatusMovedPermanently))

	home.Register(router, gamesRepo)
	gr.Register(router, gamesRepo)

	log.Println("Server is run at port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
