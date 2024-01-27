package repository

import (
	"github.com/L1z1ng3r-sswe/instagram_clone/app/internal/model"
	"github.com/jmoiron/sqlx"
)

type User interface {
}

type Post interface {
	CreatePost(newPost model.Post)
}

type Comment interface {
}

type Repository struct {
	User
	Post
	Comment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Post: newPostRepository(db),
	}
}
