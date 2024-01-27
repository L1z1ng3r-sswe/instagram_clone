package service

import (
	"github.com/L1z1ng3r-sswe/instagram_clone/app/internal/model"
	"github.com/L1z1ng3r-sswe/instagram_clone/app/internal/repository"
)

type User interface {
}

type Post interface {
	CreatePost(post model.Post)
}

type Comments interface {
}

type Service struct {
	User
	Post
	Comments
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Post: newPostService(repo.Post),
	}
}
