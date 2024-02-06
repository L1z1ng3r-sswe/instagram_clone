package handler

import (
	"net/http"
	"strconv"

	"github.com/L1z1ng3r-sswe/instagram_clone/app/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) SignIn(ctx *gin.Context) {
	var user model.UserSignIn

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Bad Request": err.Error()})
		h.log.Err("Bad Request", err.Error(), "")
		return
	}

	err, errKey, errMsg, code, fileInfo, userId, tests := h.service.SignIn(user)
	if err != nil {
		ctx.AbortWithStatusJSON(code, gin.H{errKey: errMsg})
		h.log.Err(errKey, errMsg, fileInfo)
		return
	}

	ctx.JSON(http.StatusOK, tests)
	h.log.Inf("Signed in to account: " + strconv.Itoa(userId))
}
