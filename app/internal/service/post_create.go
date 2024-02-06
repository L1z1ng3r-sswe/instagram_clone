package service

import (
	"errors"
	"mime/multipart"
	"net/http"

	"github.com/L1z1ng3r-sswe/instagram_clone/app/internal/model"
)

func (service *postService) CreatePost(file *multipart.FileHeader, files []*multipart.FileHeader, post model.CreatePost) (error, string, string, int, string, string, []string, int) {
	if len(post.Description) < 4 {
		return errors.New("Validation is not allowed"), "Bad Request", "Validation is not allowed", http.StatusBadRequest, getFileInfo("post_create.go"), "", nil, 0
	}

	// handle images
	var imagesURL []string
	var imagesPath []string
	for _, file := range files {
		err, errKey, errMsg, code, fileInfo, imageURL, imagePath := service.imageHandler(file, post.CreatedBy)
		if err != nil {
			return err, errKey, errMsg, code, fileInfo, "", nil, 0
		}

		imagesURL = append(imagesURL, imageURL)
		imagesPath = append(imagesPath, imagePath)
	}

	// handle main-image
	err, errKey, errMsg, code, fileInfo, imageURL, imagePath := service.imageHandler(file, post.CreatedBy)
	if err != nil {
		return err, errKey, errMsg, code, fileInfo, "", nil, 0
	}

	err, errKey, errMsg, code, fileInfo, postId := service.repo.CreatePost(imageURL, imagesURL, post)
	if err != nil {
		return err, errKey, errMsg, code, fileInfo, "", nil, 0
	}

	return nil, "", "", http.StatusOK, "", imagePath, imagesPath, postId
}
