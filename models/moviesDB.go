package models

import (
	"context"
	"database/sql"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

// Get returns one movie
func (m *DBModel) Get(id int) (*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	query := `select id, title, description, release_date, year, rating, runtime, mpaa_rating, created_at, updated_at 
				from movies
				where id = $1
`
	row := m.DB.QueryRowContext(ctx, query, id)

	var movie Movie

	err := row.Scan(
		&movie.ID,
		&movie.Title,
		&movie.Description,
		&movie.ReleaseDate,
		&movie.Year,
		&movie.Rating,
		&movie.Runtime,
		&movie.MPAARating,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	// Get genres,

	query = `
select mg.id, mg.movie_id, mg.genre_id, g.genre_name 
from movies_genres AS mg
left join genres AS g on g.id = mg.genre_id
where mg.movie_id = $1
`

	rows, err := m.DB.QueryContext(ctx, query, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var genres []string

	for rows.Next() {
		var mg MovieGenre

		err = rows.Scan(
			&mg.ID,
			&mg.MovieID,
			&mg.GenreID,
			&mg.Genre.GenreName,
		)

		if err != nil {
			return nil, err
		}

		genres = append(genres, mg.Genre.GenreName)
	}

	movie.MovieGenre = genres

	return &movie, nil
}

// All returns all movies
func (m *DBModel) All(ID int) (*Movie, error) {
	return nil, nil
}
