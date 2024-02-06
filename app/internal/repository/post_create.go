package repository

import (
	"net/http"

	"github.com/L1z1ng3r-sswe/instagram_clone/app/internal/model"
	"github.com/lib/pq"
)

func (repo *postRepository) CreatePost(imageURL string, imagesURL []string, post model.CreatePost) (error, string, string, int, string, int) {
	stmt, err := repo.db.Preparex("INSERT INTO posts (description, main_image, images, created_by) VALUES($1, $2, $3, $4) RETURNING id")
	if err != nil {
		return err, "Internal Server Error", err.Error(), http.StatusInternalServerError, getFileInfo("post_create.go"), 0
	}

	imagesURL_interface := pq.Array(imagesURL)

	var postId int
	if err := stmt.Get(&postId, post.Description, imageURL, imagesURL_interface, post.CreatedBy); err != nil {
		return err, "Internal Server Error", err.Error(), http.StatusInternalServerError, getFileInfo("post_create.go"), 0
	}

	return err, "", "", http.StatusOK, getFileInfo("post_create.go"), postId
}
