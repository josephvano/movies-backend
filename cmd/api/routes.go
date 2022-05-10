package main

import "github.com/julienschmidt/httprouter"

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc("GET", "/status", app.statusHandler)

	router.HandlerFunc("GET", "/v1/movies/:id", app.getOneMovie)
	router.HandlerFunc("GET", "/v1/movies", app.getOneMovie)

	return router
}