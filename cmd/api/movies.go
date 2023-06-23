package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/k1nho/letsgo/internal/data"
	"github.com/k1nho/letsgo/internal/validator"
)

// createMovieHandler: create a Movie given title, year, runtime, genres (JSON)
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	m := &data.Movie{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: input.Runtime,
		Genres:  input.Genres,
	}

	v := validator.New()

	if data.ValidateMovie(v, m); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Movies.Insert(m)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/movies/%d", m.ID))

	err = app.WriteJson(w, http.StatusCreated, envelope{"movie": m}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

// showMovieHandler: get a specific movie given id in path (JSON)
func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(w, r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movie, err := app.models.Movies.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.WriteJson(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// updateMovieHandler: update a movie given an id in path (update: title, year, runtime, genres) (JSON)
func (app *application) updateMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(w, r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movie, err := app.models.Movies.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Title   *string       `json:"title"`
		Year    *int32        `json:"year"`
		Runtime *data.Runtime `json:"runtime"`
		Genres  []string      `json:"genres"`
	}
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Title != nil {
		movie.Title = *input.Title
	}
	if input.Year != nil {

		movie.Year = *input.Year
	}
	if input.Runtime != nil {

		movie.Runtime = *input.Runtime
	}
	if input.Genres != nil {

		movie.Genres = input.Genres
	}

	v := validator.New()

	if data.ValidateMovie(v, movie); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Movies.Update(movie)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.WriteJson(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// deleteMovieHandler: Delete a movie given id in param
func (app *application) deleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(w, r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Movies.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.WriteJson(w, http.StatusOK, envelope{"message": "movie successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
