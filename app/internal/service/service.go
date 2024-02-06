package service

import (
	"mime/multipart"
	"runtime"
	"strconv"

	"github.com/L1z1ng3r-sswe/instagram_clone/app/internal/model"
	"github.com/L1z1ng3r-sswe/instagram_clone/app/internal/repository"
)

type User interface {
	SignUp(user model.UserSignUp, avatar *multipart.FileHeader) (error, string, string, int, string, string, model.Tokens, int)
	SignIn(user model.UserSignIn) (error, string, string, int, string, int, model.Tokens)
}

type Post interface {
	CreatePost(file *multipart.FileHeader, files []*multipart.FileHeader, post model.CreatePost) (error, string, string, int, string, string, []string, int)
	GetPosts() (error, string, string, int, string, []model.Post)
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
		User: newUserService(repo.User),
		Post: newPostService(repo.Post),
	}
}

func getFileInfo(fileName string) string {
	_, _, line, _ := runtime.Caller(1)
	return "app/internal/service/" + fileName + " line: " + strconv.Itoa(line)
}
