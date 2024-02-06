package repository

import (
	"net/http"

	"github.com/L1z1ng3r-sswe/instagram_clone/app/internal/model"
	"github.com/lib/pq" // Import the pq package for handling PostgreSQL array types
)

func (repo *postRepository) GetPosts() (error, string, string, int, string, []model.Post) {
	var posts []model.Post

	stmt, err := repo.db.Preparex("SELECT id, description, created_by, main_image, images FROM posts")
	defer stmt.Close()
	if err != nil {
		return err, "Internal Server Error", err.Error(), http.StatusInternalServerError, getFileInfo("post_get.go"), nil
	}

	rows, err := stmt.Queryx()
	if err != nil {
		return err, "Internal Server Error", err.Error(), http.StatusInternalServerError, getFileInfo("post_get.go"), nil
	}
	defer rows.Close()

	for rows.Next() {
		var post model.Post
		var images pq.StringArray // Using pq.StringArray to handle PostgreSQL array types

		err := rows.Scan(&post.Id, &post.Description, &post.CreatedBy, &post.MainImage, &images)
		if err != nil {
			return err, "Internal Server Error", err.Error(), http.StatusInternalServerError, getFileInfo("post_get.go"), nil
		}

		// Converting pq.StringArray to []string
		post.Images = make([]string, len(images))
		for i, img := range images {
			post.Images[i] = img
		}

		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return err, "Internal Server Error", err.Error(), http.StatusInternalServerError, getFileInfo("post_get.go"), nil
	}

	return nil, "", "", http.StatusOK, "", posts
}
