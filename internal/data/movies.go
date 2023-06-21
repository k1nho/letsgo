package data

import (
	"time"

	"github.com/k1nho/letsgo/internal/validator"
)

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	Year      int32     `json:"year"`
	Runtime   Runtime   `json:"runtime,omitempty"`
	Genres    []string  `json:"genres"`
	Version   int32     `json:"version"`
}

/* Validator Contraints
   -- Title: must not be empty and must be less than or equal to 500 bytes long
   -- Year: must not be 0, and must be greater than or equal to 1888 and cannot be in the future
   -- Runtime: must not be 0 and has to be greater than 0
   -- Genres: must not be nil and has to contain at least 1 and at most 5, also all the genres provided must be unique
*/

func ValidateMovie(v *validator.Validator, m *Movie) {
	v.Check(m.Title != "", "title", "must be provided")
	v.Check(len(m.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(m.Year != 0, "year", "must be provided")
	v.Check(m.Year >= 1888, "year", "must be greater than 1888")
	v.Check(m.Year < int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(m.Runtime != 0, "runtime", "must be provided")
	v.Check(m.Runtime > 0, "runtime", "must be positive")

	v.Check(m.Genres != nil, "genres", "must be provided")
	v.Check(len(m.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(m.Genres) <= 5, "genres", "must not exceeed 5 genres")
	v.Check(validator.Unique(m.Genres), "genres", "must not contain duplicates")

}
