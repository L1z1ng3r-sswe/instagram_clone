package service

import (
	"net/http"

	"github.com/L1z1ng3r-sswe/instagram_clone/app/internal/model"
)

func (service *postService) GetPosts() (error, string, string, int, string, []model.Post) {
	err, errKey, errMsg, code, fileInfo, posts := service.repo.GetPosts()
	if err != nil {
		return err, errKey, errMsg, code, fileInfo, nil
	}

	return nil, "", "", http.StatusOK, "", posts
}
