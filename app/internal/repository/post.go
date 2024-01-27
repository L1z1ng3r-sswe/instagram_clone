package repository

import "github.com/jmoiron/sqlx"

type postRepository struct {
	db *sqlx.DB
}

func newPostRepository(db *sqlx.DB) *postRepository {
	return &postRepository{
		db: db,
	}
}
