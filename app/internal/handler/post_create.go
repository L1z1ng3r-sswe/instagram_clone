package handler

import (
	"fmt"
	"net/http"

	"github.com/L1z1ng3r-sswe/instagram_clone/app/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePost(ctx *gin.Context) {
	var post model.CreatePost

	if err := ctx.ShouldBind(&post); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Bad Request": err.Error()})
		h.log.Err("Bad Request", err.Error(), "")
		return
	}

	file, err := ctx.FormFile("main_image")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Bad Request": err.Error()})
		h.log.Err("Bad Request", err.Error(), "")
		return
	}

	files := ctx.Request.MultipartForm.File["images"]

	err, errKey, errMsg, code, fileInfo, imagePath, imagesPath, postId := h.service.Post.CreatePost(file, files, post)
	if err != nil {
		ctx.AbortWithStatusJSON(code, gin.H{errKey: errMsg})
		h.log.Err(errKey, errMsg, fileInfo)
		return
	}

	for _, imagePath := range imagesPath {
		if err := ctx.SaveUploadedFile(file, imagePath); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Internal Server Error": err.Error()})
			h.log.Err("Internal Server Error", err.Error(), "")
			return
		}
	}

	if err := ctx.SaveUploadedFile(file, imagePath); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Internal Server Error": err.Error()})
		h.log.Err("Internal Server Error", err.Error(), "")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Created a new post: ": postId})
	h.log.Inf(fmt.Sprint("Created a new post: ", postId))
}
