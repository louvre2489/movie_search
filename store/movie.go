package store

import (
	"context"

	"github.com/louvre2489/movie_search/entity"
)

func (r *Repository) ListTasks(
	ctx context.Context,
	db Queryer,
) (entity.Movies, error) {
	movies := entity.Movies{}
	sql := `SELECT * FROM movie;`

	if err := db.SelectContext(ctx, &movies, sql); err != nil {
		return nil, err
	}

	return movies, nil
}
