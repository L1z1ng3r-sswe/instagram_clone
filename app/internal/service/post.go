package service

import (
	"github.com/L1z1ng3r-sswe/instagram_clone/app/internal/repository"
)

type postService struct {
	repo repository.Post
}

func newPostService(repo repository.Post) *postService {
	return &postService{
		repo: repo,
	}
}
