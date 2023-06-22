package data

import (
	"database/sql"
	"errors"
	"time"

	"github.com/k1nho/letsgo/internal/validator"
	"github.com/lib/pq"
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

type MovieModel struct {
	DB *sql.DB
}

func (m MovieModel) Get(id int64) (*Movie, error) {

	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
        SELECT id, created_at, title, year, runtime, genres, version
        FROM movies
        WHERE id=$1
    `

	var movie Movie

	err := m.DB.QueryRow(query, id).Scan(&movie.ID, &movie.CreatedAt, &movie.Title, &movie.Year, &movie.Runtime, pq.Array(&movie.Genres), &movie.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &movie, nil
}

func (m MovieModel) Insert(movie *Movie) error {
	query := `
        INSERT INTO movies(title, year, runtime, genres)
        VALUES($1, $2, $3, $4)
        RETURNING id, created_at, version
    `

	args := []interface{}{movie.Title, movie.Year, movie.Runtime, pq.Array(movie.Genres)}

	return m.DB.QueryRow(query, args...).Scan(&movie.ID, &movie.CreatedAt, &movie.Version)
}

func (m MovieModel) Update(movie *Movie) error {

	query := `
        UPDATE movies
        SET title=$1, year=$2, runtime=$3, genres=$4, version=version+1
        WHERE id=$5
        RETURNING version
    `

	args := []interface{}{movie.Title, movie.Year, movie.Runtime, pq.Array(movie.Genres), movie.ID}

	return m.DB.QueryRow(query, args...).Scan(&movie.Version)

}

func (m MovieModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
        DELETE FROM movies
        WHERE id=$1
    `

	result, err := m.DB.Exec(query, id)
	if err != nil {
		return err
	}

	nRows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if nRows == 0 {
		return ErrRecordNotFound
	}

	return nil
}
