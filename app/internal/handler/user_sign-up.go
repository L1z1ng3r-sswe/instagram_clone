package handler

import (
	"net/http"
	"strconv"

	"github.com/L1z1ng3r-sswe/instagram_clone/app/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(ctx *gin.Context) {
	var user model.UserSignUp
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Bad Request": err.Error()})
		h.log.Err("Bad Request", err.Error(), "")
		return
	}

	avatar, err := ctx.FormFile("avatar")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Bad Request": err.Error()})
		h.log.Err("Bad Request", err.Error(), "")
		return
	}

	err, errKey, errMsg, code, fileInfo, avatarPath, tokens, userId := h.service.SignUp(user, avatar)
	if err != nil {
		ctx.AbortWithStatusJSON(code, gin.H{errKey: errMsg})
		h.log.Err(errKey, errMsg, fileInfo)
		return
	}

	if err := ctx.SaveUploadedFile(avatar, avatarPath); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Internal Server Error": err.Error()})
		h.log.Err("Internal Server Error", err.Error(), "")
		return
	}

	ctx.JSON(http.StatusOK, tokens)
	h.log.Inf("Signed up a new user: " + strconv.Itoa(userId))
}
