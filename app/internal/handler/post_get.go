package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetPosts(ctx *gin.Context) {
	err, errKey, errMsg, code, fileInfo, posts := h.service.GetPosts()
	if err != nil {
		ctx.AbortWithStatusJSON(code, gin.H{errKey: errMsg})
		h.log.Err(errKey, errMsg, fileInfo)
		return
	}

	ctx.JSON(http.StatusOK, posts)
}
