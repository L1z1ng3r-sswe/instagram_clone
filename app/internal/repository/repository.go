package repository

import (
	"runtime"
	"strconv"

	"github.com/L1z1ng3r-sswe/instagram_clone/app/internal/model"
	"github.com/jmoiron/sqlx"
)

type User interface {
	SignUp(user model.UserSignUp, imageURL string) (error, string, string, int, string, int)
	SignIn(user model.UserSignIn) (error, string, string, int, string, string, int)
}

type Post interface {
	CreatePost(imagePath string, imagesPath []string, newPost model.CreatePost) (error, string, string, int, string, int)
	GetPosts() (error, string, string, int, string, []model.Post)
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
		User: NewUserRepository(db),
	}
}

func getFileInfo(fileName string) string {
	_, _, line, _ := runtime.Caller(1)
	return "app/internal/repository/" + fileName + " line: " + strconv.Itoa(line)
}
