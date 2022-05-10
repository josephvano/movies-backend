package main

import (
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)

		return
	}

	app.logger.Println("id is", id)

	movie, err := app.models.DB.Get(id)

	if err != nil {
		app.logger.Print(errors.New("could not get movie with id " + strconv.Itoa(id)))
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, movie, "movie")

	if err != nil {
		app.logger.Println(err)
	}
}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {

}
