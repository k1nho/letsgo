package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/k1nho/letsgo/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Created a new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(w, r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "The Matrix",
		Year:      1999,
		Runtime:   201,
		Genres:    []string{"psychological", "Avant-Garde", "Action"},
		Version:   1,
	}

	err = app.WriteJson(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.logger.Print(err)
		http.Error(w, "Could not find the movie", http.StatusInternalServerError)
	}
}