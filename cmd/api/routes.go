package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc("GET", "/status", app.statusHandler)

	router.HandlerFunc("GET", "/v1/movies/:id", app.getOneMovie)
	router.HandlerFunc("GET", "/v1/movies", app.getAllMovies)

	// genre routes
	router.HandlerFunc("GET", "/v1/genres", app.getAllGenres)

	return app.enableCORS(router)
}