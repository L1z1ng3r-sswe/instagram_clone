package service

import (
	"errors"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

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

// !================================ image handler =============================================
func (service *postService) imageHandler(file *multipart.FileHeader, createdBy int) (error, string, string, int, string, string, string) {
	if !service.imageValidation(file) {
		return errors.New("Incorrect file type"), "Bad Request", "Incorrect file type", http.StatusBadRequest, getFileInfo("post_create.go"), "", ""
	}

	err, imagePath := service.saveImage(file, strconv.Itoa(createdBy))
	if err != nil {
		return err, "Internal Server Error", err.Error(), http.StatusBadRequest, getFileInfo("post_create.go"), "", ""
	}

	imageURL := service.constructURL(imagePath)
	return nil, "", "", 200, "", imageURL, imagePath
}

func (service *postService) imageValidation(file *multipart.FileHeader) bool {
	extension := strings.ToLower(filepath.Ext(file.Filename))
	if extension != ".jpg" && extension != ".png" && extension != ".jpeg" {
		return false
	}

	return true
}

func (service *postService) saveImage(file *multipart.FileHeader, userId string) (error, string) {
	folderPath := "app/pkg/storage/post_image/" + userId
	if err := os.MkdirAll(folderPath, 0755); err != nil {
		return err, ""
	}

	filePath := filepath.Join(folderPath, file.Filename)
	return nil, filePath
}

func (service *postService) constructURL(filePath string) string {
	filePath = strings.TrimPrefix(filePath, "app/pkg/storage")
	apiPort := os.Getenv("API_PORT")

	return "http://localhost:" + apiPort + filePath
}
